import '@testing-library/jest-dom';
import { render } from '@testing-library/react';
import { TileEntity } from '@/types';
import TileEntityList from './TileEntityList';

const testTileData: TileEntity[] = [
  {
    name: `Abyssal Sire`,
    safeURI: `AbyssalSire`,
    thumbnail: `https://oldschool.runescape.wiki/images/thumb/Abyssal_Sire.png/200px-Abyssal_Sire.png`,
    wiki: `https://oldschool.runescape.wiki/w/Abyssal_Sire`,
    tags: [`slayer`, `boss`],
    recommendedGuideVideoId: `wnZJl9driUs`,
    source: `https://www.youtube.com/watch?v=wnZJl9driUs`,
    tiles: [
      {
        regionId: 11850,
        regionX: 25,
        regionY: 34,
        z: 0,
        color: `#FFFFFF00`,
      },
    ],
  },
  {
    name: `Alchemical Hydra`,
    safeURI: `AlchemicalHydra`,
    thumbnail: `https://oldschool.runescape.wiki/images/thumb/Alchemical_Hydra.png/200px-Alchemical_Hydra.png`,
    wiki: `https://oldschool.runescape.wiki/w/Alchemical_Hydra`,
    tags: [`slayer`, `boss`],
    recommendedGuideVideoId: `7ehTfsD--gM`,
    source: `https://www.youtube.com/watch?v=7ehTfsD--gM`,
    tiles: [
      {
        regionId: 5536,
        regionX: 27,
        regionY: 38,
        z: 0,
        color: `#FFFFFF00`,
      },
      {
        regionId: 5537,
        regionX: 27,
        regionY: 38,
        z: 0,
        color: `#FFFFFF00`,
      },
    ],
  },
];

describe(`TileEntityList`, () => {
  it(`should map entities correctly`, () => {
    const { getByText } = render(<TileEntityList list={testTileData} />);
    expect(getByText(`Abyssal Sire`)).toBeInTheDocument();
    expect(getByText(`Alchemical Hydra`)).toBeInTheDocument();
    expect(getByText(`1 tile markers`)).toBeInTheDocument();
    expect(getByText(`2 tile markers`)).toBeInTheDocument();
  });

  // it(`should scroll to the top on search`, () => {});
});
