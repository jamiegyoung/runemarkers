export type Tile = {
  regionId: number;
  regionX: number;
  regionY: number;
  z: number;
  color: string;
};

export type TileEntity = {
  name: string;
  tags: string[];
  tiles: Tile[];
  thumbnail: string;
  wiki: string;
  source?: string;
  recommendedGuide?: string;
};
