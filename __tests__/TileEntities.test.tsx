import '@testing-library/jest-dom';
import { act, render } from '@testing-library/react';
import { TileEntity } from '@/types';
import TileEntities, { ViewFormat } from '@/components/organisms/TileEntities';

const testTileData: TileEntity[] = [
  {
    name: `Abyssal Sire`,
    safeURI: `AbyssalSire`,
    thumbnail: `https://oldschool.runescape.wiki/images/thumb/Abyssal_Sire.png/200px-Abyssal_Sire.png`,
    wiki: `https://oldschool.runescape.wiki/w/Abyssal_Sire`,
    tags: [`slayer`, `boss`],
    recommendedGuideVideoId: `wnZJl9driUs`,
    source: {
      name: `AsukaYen OSRS - OSRS Abyssal Sire Guide [2021]`,
      link: `https://www.youtube.com/watch?v=wnZJl9driUs`,
    },
    tiles: [
      {
        regionId: 11850,
        regionX: 25,
        regionY: 34,
        z: 0,
        color: `#FFFFFF00`,
      },
    ],
    fullName: `Abyssal Sire`,
    fullAltName: ``,
  },
  {
    name: `Alchemical Hydra`,
    safeURI: `AlchemicalHydra`,
    thumbnail: `https://oldschool.runescape.wiki/images/thumb/Alchemical_Hydra.png/200px-Alchemical_Hydra.png`,
    wiki: `https://oldschool.runescape.wiki/w/Alchemical_Hydra`,
    tags: [`slayer`, `boss`],
    recommendedGuideVideoId: `7ehTfsD--gM`,
    source: {
      name: `AsukaYen OSRS - OSRS Abyssal Sire Guide [2021]`,
      link: `https://www.youtube.com/watch?v=wnZJl9driUs`,
    },
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
    fullName: `Alchemical Hydra`,
    fullAltName: ``,
  },
];

describe(`TileEntites`, () => {
  beforeEach(() => {
    window.localStorage.clear();
  });

  it(`should map entities correctly`, () => {
    const { getByText } = render(<TileEntities list={testTileData} />);
    expect(getByText(`Abyssal Sire`)).toBeInTheDocument();
    expect(getByText(`Alchemical Hydra`)).toBeInTheDocument();
    expect(getByText(`1 tile markers`)).toBeInTheDocument();
    expect(getByText(`2 tile markers`)).toBeInTheDocument();
  });

  it(`should display a message if no entities are found`, () => {
    const { getByText } = render(<TileEntities list={[]} />);
    expect(getByText(`No results found`)).toBeInTheDocument();
  });

  it(`should be able to switch between list and grid view`, () => {
    const { container } = render(<TileEntities list={testTileData} />);
    const listContainer = container.querySelector(`.innerListContainer`);

    expect(listContainer).toBeInTheDocument();
    const gridButton = container.querySelector(`.button`) as HTMLButtonElement;
    if (!gridButton) {
      throw new Error(`No button found`);
    }

    act(() => {
      gridButton.click();
    });

    const gridContainer = container.querySelector(`.innerGridContainer`);
    expect(gridContainer).toBeInTheDocument();
    expect(listContainer).not.toBeInTheDocument();
  });

  it(`should set the view to a list by default`, () => {
    const { container } = render(<TileEntities list={testTileData} />);
    // wait for the list to render
    const listContainer = container.querySelector(`.innerListContainer`);
    expect(listContainer).toBeInTheDocument();
    // output render
  });

  it(`should set the view to a list if the window is smaller than 600px`, () => {
    const { container } = render(<TileEntities list={testTileData} />);
    const button = container.querySelector(`.button`) as HTMLButtonElement;
    if (!button) {
      throw new Error(`No button found`);
    }

    act(() => {
      button.click();
    });

    const gridContainer = container.querySelector(`.innerGridContainer`);
    expect(gridContainer).toBeInTheDocument();

    Object.defineProperty(window, `innerWidth`, {
      writable: true,
      configurable: true,
      value: 400,
    });

    act(() => {
      window.dispatchEvent(new Event(`resize`));
    });
    const listContainer = container.querySelector(`.innerListContainer`);

    expect(gridContainer).not.toBeInTheDocument();
    expect(listContainer).toBeInTheDocument();
  });

  it(`should save the view to local storage`, () => {
    jest.spyOn(localStorage, `setItem`);
    localStorage.setItem = jest.fn();

    jest.spyOn(Storage.prototype, `setItem`);
    Storage.prototype.setItem = jest.fn();

    const { container } = render(<TileEntities list={testTileData} />);
    const gridButton = container.querySelector(`.button`) as HTMLButtonElement;
    if (!gridButton) {
      throw new Error(`No button found`);
    }
    act(() => {
      gridButton.click();
    });
    expect(localStorage.setItem).toBeCalledWith(
      `viewFormat`,
      `${ViewFormat.Grid}`,
    );
  });

  it(`should load a view from local storage`, () => {
    jest.spyOn(localStorage, `getItem`);
    localStorage.getItem = jest.fn();

    jest.spyOn(Storage.prototype, `getItem`);
    Storage.prototype.getItem = jest.fn();

    render(<TileEntities list={testTileData} />);
    expect(localStorage.getItem).toBeCalledWith(`viewFormat`);
  });
});
