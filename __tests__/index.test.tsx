import { getTileData } from '@/api/tiles';
import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import Home, { getStaticProps } from '@/pages/index';
import { MappedTileEntity } from '@/types';

describe(`Home`, () => {
  it(`should search correctly`, () => {
    const data = getTileData();
    const mappedData: MappedTileEntity[] = data.map((e) => ({
      ...e,
      fullName: `${e.name} ${e.subcategory}`,
      fullAltName: `${e.altName} ${e.subcategory}`,
    }));
    const { container, getByPlaceholderText } = render(
      <Home tileData={mappedData} />,
    );
    const input = getByPlaceholderText(`type here to search`);
    act(() => {
      fireEvent.change(input, { target: { value: `abyssal sire` } });
    });
    expect(container).toMatchSnapshot();
  });

  it(`should generate static props`, async () => {
    const tileData = getTileData();
    const mappedTileData: MappedTileEntity[] = tileData.map((e) => ({
      ...e,
      fullName: `${e.name} ${e.subcategory}`,
      fullAltName: `${e.altName} ${e.subcategory}`,
    }));
    const { props } = await getStaticProps();
    expect(props.tileData).toEqual(mappedTileData);
  });
});
