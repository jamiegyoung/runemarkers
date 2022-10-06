import '@testing-library/jest-dom';
import { getTileData } from '@/api/tiles';

describe(`Tiles`, () => {
  it(`should be in the correct order`, () => {
    const data = getTileData();
    const sorted = [...data];
    sorted.sort((a, b) =>
      a.subcategory
        ? `${a.name} (${a.subcategory})`
            .toLowerCase()
            .localeCompare(`${b.name} (${b.subcategory})`.toLowerCase())
        : a.name.localeCompare(b.name),
    );

    expect(data).toEqual(sorted);
  });
});
