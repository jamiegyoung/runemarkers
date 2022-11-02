import { TileEntity } from '@/types';
import fs from 'fs';

export const getTileData: () => TileEntity[] = () =>
  fs
    .readdirSync(`./src/tiles`, {
      encoding: `utf-8`,
      withFileTypes: true,
    })
    .filter((file) => file.isFile() && file.name.endsWith(`.json`))
    .map((file) => ({
      safeURI: file.name.replace(`.json`, ``),
      ...JSON.parse(
        fs.readFileSync(`./src/tiles/${file.name}`, { encoding: `utf-8` }),
      ),
    }))
    .map((tile) => ({
      ...tile,
      fullName: `${tile.name} ${tile.subcategory}`,
      fullAltName: `${tile.altName} ${tile.subcategory}`,
    }));
