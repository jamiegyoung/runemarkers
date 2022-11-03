import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import { TileEntity } from '@/types';
import TileEntityCardCompact from '@/components/molecules/TileEntityCardCompact';

const entity: TileEntity = {
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
};

const TestTileEntityCard = () => <TileEntityCardCompact entity={entity} />;

const TestTileEntityCardSubcategory = () => (
  <TileEntityCardCompact
    entity={{ ...entity, subcategory: `Test Subcategory` }}
  />
);

window.prompt = jest.fn();

describe(`TileEntityCard`, () => {
  it(`should render entity correctly`, () => {
    const { container, getByText } = render(<TestTileEntityCard />);
    expect(getByText(entity.name)).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it(`should handle copy click correctly`, () => {
    render(<TestTileEntityCard />);
    const copyButton = document.querySelector(`.button`);
    if (!copyButton) {
      throw new Error(`Copy button not found`);
    }
    act(() => {
      fireEvent.click(copyButton);
    });
    const animBar = document.querySelector(`.animBarActive`);
    if (!animBar) {
      throw new Error(`Animation bar not found`);
    }
    expect(animBar).toBeInTheDocument();
    expect(window.prompt).toHaveBeenCalledTimes(1);
  });

  it(`should call onload for the image correctly`, async () => {
    const callback = jest.fn();
    const { container } = render(<TestTileEntityCard />);
    const image = container.querySelector(`img`);
    if (!image) {
      throw new Error(`Image not found`);
    }
    image.src = ``;
    await act(async () => {
      expect(image.style.opacity).toBe(`0`);
      image.onload = callback;
      image.src = entity.thumbnail;
      fireEvent.load(image);
      expect(callback).toHaveBeenCalledTimes(1);
      // wait for the animation to finish
      setTimeout(() => {
        expect(image.style.opacity).toBe(`1`);
      }, 1000);
    });
  });

  it(`should render subcategory correctly`, () => {
    const { container, getByText } = render(<TestTileEntityCardSubcategory />);
    expect(getByText(entity.name)).toBeInTheDocument();
    expect(getByText(`(Test Subcategory)`)).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });
});
