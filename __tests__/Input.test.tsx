import { SearchContext } from '@/pages';
import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import { Dispatch, SetStateAction } from 'react';
import Input from '@/components/atoms/Input';

const testInputWithContext = (
  searchVal: string | undefined,
  setSearchVal: Dispatch<SetStateAction<string>>,
) => (
  <SearchContext.Provider value={[searchVal, setSearchVal]}>
    <div>
      <Input />
      <button></button>
    </div>
  </SearchContext.Provider>
);
describe(`Input`, () => {
  it(`should render correctly`, () => {
    const { container } = render(<Input />);
    expect(container).toMatchSnapshot();
  });

  it(`should expand on focus`, () => {
    const mockSetSearchVal = jest.fn();

    jest.mock(`react`, () => ({
      useState: (initial: string) => [initial, mockSetSearchVal],
    }));

    const { container } = render(testInputWithContext(``, mockSetSearchVal));

    const input = container.querySelector(`input`);
    if (!input) {
      throw new Error(`No input found`);
    }
    act(() => {
      fireEvent.focus(input);
    });
    expect(input.placeholder).toBe(``);
  });

  it(`should shrink on blur`, () => {
    const mockSetSearchVal = jest.fn();

    jest.mock(`react`, () => ({
      useState: (initial: string) => [initial, mockSetSearchVal],
    }));

    const { container } = render(testInputWithContext(``, mockSetSearchVal));

    const input = container.querySelector(`input`);
    if (!input) {
      throw new Error(`No input found`);
    }
    act(() => {
      fireEvent.focus(input);
    });
    expect(input.placeholder).toBe(``);
    act(() => {
      input.value = ``;
      fireEvent.blur(input);
    });
    expect(input.placeholder).toBe(`type here to search`);
  });

  it(`should update the search value on change`, () => {
    const mockSetSearchVal = jest.fn();

    jest.mock(`react`, () => ({
      useState: (initial: string) => [initial, mockSetSearchVal],
    }));

    const { container } = render(testInputWithContext(``, mockSetSearchVal));

    const input = container.querySelector(`input`);
    if (!input) {
      throw new Error(`No input found`);
    }
    act(() => {
      fireEvent.focus(input);
    });
    expect(input.placeholder).toBe(``);
    act(() => {
      fireEvent.change(input, { target: { value: `test` } });
    });
    expect(mockSetSearchVal).toHaveBeenCalledTimes(1);
    expect(mockSetSearchVal).toHaveBeenCalledWith(`test`);
  });
});
