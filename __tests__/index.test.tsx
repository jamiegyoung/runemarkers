import { getTileData } from '@/api/tiles';
import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import Home, { getStaticProps } from '@/pages/index';

describe(`Home`, () => {
  it(`should search correctly`, () => {
    const data = getTileData();
    const { container, getByPlaceholderText } = render(
      <Home tileData={data} />,
    );
    const input = getByPlaceholderText(`search tiles`);
    act(() => {
      fireEvent.change(input, { target: { value: `abyssal sire` } });
    });
    expect(container).toMatchSnapshot();
  });

  it(`should generate static props`, () => {
    const tileData = getTileData();
    const { props } = getStaticProps();
    expect(props.tileData).toEqual(tileData);
  });
});
