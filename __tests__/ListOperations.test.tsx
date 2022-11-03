import ListOperations from '@/components/molecules/ListOperations';
import { ViewFormat } from '@/components/organisms/TileEntities';
import '@testing-library/jest-dom';
import { act, render } from '@testing-library/react';

describe(`ListOperations`, () => {
  it(`should render correctly with the list format`, () => {
    const callback = jest.fn();
    const { container } = render(
      <ListOperations viewFormat={ViewFormat.List} setViewFormat={callback} />,
    );
    expect(container).toMatchSnapshot();
  });

  it(`should render correctly with the grid format`, () => {
    const callback = jest.fn();
    const { container } = render(
      <ListOperations viewFormat={ViewFormat.Grid} setViewFormat={callback} />,
    );
    expect(container).toMatchSnapshot();
  });

  it(`should call the callback when the list button is clicked`, () => {
    const callback = jest.fn();
    const { container } = render(
      <ListOperations viewFormat={ViewFormat.List} setViewFormat={callback} />,
    );
    const button = container.querySelector(`.button`) as HTMLButtonElement;
    if (!button) {
      throw new Error(`No button found`);
    }
    act(() => {
      button.click();
    });
    expect(callback).toBeCalledWith(ViewFormat.Grid);
  });
});
