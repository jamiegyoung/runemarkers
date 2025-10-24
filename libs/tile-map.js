const regionToWorld = (regionId, regionX, regionY) => {
	const worldX = ((regionId >>> 8) << 6) + regionX;
	const worldY = ((regionId & 0xff) << 6) + regionY;
	return { x: worldX, y: worldY };
};

const toLatLng = (x, y) => [y, x];

const fromLatLng = (latLng) => ({
	x: parseInt(latLng.lng),
	y: parseInt(latLng.lat)
});

const parseHexColor = (hex) => {
	const a = parseInt(hex.substring(0, 2), 16) / 255;
	const r = parseInt(hex.substring(2, 4), 16);
	const g = parseInt(hex.substring(4, 6), 16);
	const b = parseInt(hex.substring(6, 8), 16);
	return `rgba(${r},${g},${b},${a})`;
};

const parseIntColor = (num) => {
	const value = num >>> 0;
	const b = value & 0xff;
	const g = (value & 0xff00) >>> 8;
	const r = (value & 0xff0000) >>> 16;
	const a = ((value & 0xff000000) >>> 24) / 255;
	return `rgba(${r},${g},${b},${a})`;
};

const colorToRGBA = (color) => {
	if (typeof color === 'string' && color.startsWith('#')) {
		return parseHexColor(color.substring(1));
	}

	if (typeof color === 'object' && color.value !== undefined) {
		return parseIntColor(color.value);
	}

	if (typeof color === 'number') {
		return parseIntColor(color);
	}

	return 'rgba(255,255,255,0.5)';
};

const transformMarkers = (tiles) => {
	return tiles.map(tile => {
		const { x, y } = regionToWorld(tile.regionId, tile.regionX, tile.regionY);
		return {
			x,
			y,
			z: tile.z || 0,
			color: colorToRGBA(tile.color),
			label: tile.label || ''
		};
	});
};

class OSRSTileLayer extends L.TileLayer {
	constructor(options) {
		const defaultOptions = {
			minZoom: -4,
			maxZoom: 8,
			minNativeZoom: 0,
			maxNativeZoom: 2,
			tileSize: 256,
			errorTileUrl: 'https://raw.githubusercontent.com/mejrs/mejrs.github.io/master/layers/alpha_pixel.png'
		};
		super('', { ...defaultOptions, ...options });
		this.currentPlane = options.plane || 0;
	}

	getTileUrl(coords) {
		const zoom = coords.z;
		const x = coords.x;
		const y = -(1 + coords.y);
		const plane = this.currentPlane;
		return `https://raw.githubusercontent.com/mejrs/layers_osrs/master/mapsquares/-1/${zoom}/${plane}_${x}_${y}.png`;
	}

	setPlane(plane) {
		this.currentPlane = plane;
		this.redraw();
	}
}

class TileMap {
	constructor(containerId, tiles, options = {}) {
		this.containerId = containerId;
		this.tiles = tiles;
		this.options = options;
		this.currentPlane = 0;
		this.markers = [];
		this.hoverRect = null;
		this.initialBounds = null;

		this.init();
	}

	init() {
		const transformedTiles = transformMarkers(this.tiles);
		this.markers = transformedTiles;
		this.initialBounds = this.calculateBounds(transformedTiles);
		this.currentPlane = this.getMostCommonPlane();

		this.map = L.map(this.containerId, {
			crs: L.CRS.Simple,
			minZoom: -4,
			maxZoom: 8,
			zoom: 0,
			zoomControl: false,
			attributionControl: false
		});

		this.map.fitBounds([
			toLatLng(this.initialBounds.minX, this.initialBounds.minY),
			toLatLng(this.initialBounds.maxX, this.initialBounds.maxY)
		]);

		this.tileLayer = new OSRSTileLayer({ plane: this.currentPlane });
		this.tileLayer.addTo(this.map);

		this.renderMarkers(this.currentPlane);
		this.addHoverRect();
		this.addControls();
	}

	calculateBounds(tiles) {
		if (tiles.length === 0) {
			return { minX: 3200, maxX: 3250, minY: 3200, maxY: 3250 };
		}

		const xs = tiles.map(t => t.x);
		const ys = tiles.map(t => t.y);
		const padding = 4;

		return {
			minX: Math.min(...xs) - padding,
			maxX: Math.max(...xs) + padding + 1,
			minY: Math.min(...ys) - padding,
			maxY: Math.max(...ys) + padding + 1
		};
	}

	renderMarkers(plane) {
		if (this.markerLayer) {
			this.markerLayer.remove();
		}

		this.markerLayer = L.layerGroup();
		const planeMarkers = this.markers.filter(m => m.z === plane);

		planeMarkers.forEach(marker => {
			const bounds = [
				toLatLng(marker.x, marker.y),
				toLatLng(marker.x + 1, marker.y + 1)
			];

			const rect = L.rectangle(bounds, {
				color: marker.color,
				fillColor: marker.color,
				fillOpacity: 0.3,
				weight: 1,
				interactive: !!marker.label
			});

			if (marker.label) {
				rect.bindTooltip(marker.label, {
					permanent: false,
					direction: 'top',
					opacity: 1
				});
			}

			rect.addTo(this.markerLayer);
		});

		this.markerLayer.addTo(this.map);
	}

	addHoverRect() {
		const hoverStyle = {
			color: 'rgba(255, 255, 255, 0.5)',
			fillColor: 'transparent',
			weight: 1,
			interactive: false
		};

		this.hoverRect = L.rectangle([[0, 0], [1, 1]], hoverStyle);

		this.map.on('mousemove', (e) => {
			const coords = fromLatLng(e.latlng);
			const x = Math.floor(coords.x);
			const y = Math.floor(coords.y);

			this.hoverRect.setBounds([
				toLatLng(x, y),
				toLatLng(x + 1, y + 1)
			]);

			if (!this.map.hasLayer(this.hoverRect)) {
				this.hoverRect.addTo(this.map);
			}
		});

		this.map.on('mouseout', () => {
			if (this.map.hasLayer(this.hoverRect)) {
				this.hoverRect.remove();
			}
		});
	}

	createPlaneControls() {
		const control = L.control({ position: 'topright' });
		control.onAdd = () => {
			const div = L.DomUtil.create('div', 'tilemap-controls');
			div.innerHTML = `
				<div class="tilemap-control-group">
					<button class="tilemap-btn tilemap-plane-up" title="Floor Up">▲</button>
					<div class="tilemap-plane-label"><div>Floor</div><div>${this.currentPlane}</div></div>
					<button class="tilemap-btn tilemap-plane-down" title="Floor Down">▼</button>
				</div>
			`;
			L.DomEvent.disableClickPropagation(div);
			return div;
		};
		return control;
	}

	createZoomControls() {
		const control = L.control({ position: 'topleft' });
		control.onAdd = () => {
			const div = L.DomUtil.create('div', 'tilemap-controls');
			div.innerHTML = `
				<div class="tilemap-control-group">
					<button class="tilemap-btn tilemap-zoom-in" title="Zoom In">+</button>
					<button class="tilemap-btn tilemap-zoom-out" title="Zoom Out">−</button>
					<button class="tilemap-btn tilemap-reset" title="Reset View">⟲</button>
				</div>
			`;
			L.DomEvent.disableClickPropagation(div);
			return div;
		};
		return control;
	}

	addControls() {
		this.createPlaneControls().addTo(this.map);
		this.createZoomControls().addTo(this.map);
		this.attachControlHandlers();
	}

	handlePlaneUp() {
		if (this.currentPlane >= 3) return;

		this.currentPlane++;
		this.updatePlane();
	}

	handlePlaneDown() {
		if (this.currentPlane <= 0) return;

		this.currentPlane--;
		this.updatePlane();
	}

	attachControlHandlers() {
		const container = document.getElementById(this.containerId);

		container.querySelector('.tilemap-plane-up')?.addEventListener('click', () => this.handlePlaneUp());
		container.querySelector('.tilemap-plane-down')?.addEventListener('click', () => this.handlePlaneDown());
		container.querySelector('.tilemap-zoom-in')?.addEventListener('click', () => this.map.zoomIn());
		container.querySelector('.tilemap-zoom-out')?.addEventListener('click', () => this.map.zoomOut());
		container.querySelector('.tilemap-reset')?.addEventListener('click', () => this.reset());
	}

	updatePlane() {
		this.tileLayer.setPlane(this.currentPlane);
		this.renderMarkers(this.currentPlane);

		const label = document.querySelector('.tilemap-plane-label');
		if (label) {
			label.innerHTML = `<div>Floor</div><div>${this.currentPlane}</div>`;
		}
	}

	getMostCommonPlane() {
		if (!this.markers || this.markers.length === 0) return 0;

		const planeCounts = {};
		this.markers.forEach(m => {
			planeCounts[m.z] = (planeCounts[m.z] || 0) + 1;
		});
		return parseInt(Object.keys(planeCounts).sort((a, b) => planeCounts[b] - planeCounts[a])[0] || 0);
	}

	reset() {
		this.map.fitBounds([
			toLatLng(this.initialBounds.minX, this.initialBounds.minY),
			toLatLng(this.initialBounds.maxX, this.initialBounds.maxY)
		]);

		const initialPlane = this.getMostCommonPlane();

		if (this.currentPlane === initialPlane) return;

		this.currentPlane = initialPlane;
		this.updatePlane();
	}
}

window.initTileMap = (containerId, tilesJson) => {
	const tiles = typeof tilesJson === 'string' ? JSON.parse(tilesJson) : tilesJson;
	return new TileMap(containerId, tiles);
};
