export type Tile = {
  regionId: number;
  regionX: number;
  regionY: number;
  z: number;
  color: string;
  label?: string;
};

export type Source = {
  link: string;
  name: string;
  modified?: string;
};

export type TileEntity = {
  safeURI: string;
  name: string;
  subcategory?: string;
  altName?: string;
  tags: string[];
  tiles: Tile[];
  thumbnail: string;
  wiki: string;
  source?: Source;
  recommendedGuideVideoId?: string;
  fullName: string;
  fullAltName: string;
};
