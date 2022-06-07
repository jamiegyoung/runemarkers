import tileJson from '@/tiles.json';
import { TileEntity } from '@/types';

export const tileData: TileEntity[] = tileJson.map((e) => ({
  safeURI: e.name.replaceAll(`:`, `-`),
  ...e,
}));
