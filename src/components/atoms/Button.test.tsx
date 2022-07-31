import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import Button from './Button';

const callback = jest.fn();

const testingButton = (
  <Button
    onClick={() => {
      callback();
    }}
  >
    Test
  </Button>
);

const testingButtonWithClassName = (
  <Button
    onClick={() => {
      callback();
    }}
    className="test-class"
  >
    Test
  </Button>
);

describe(`Button`, () => {
  it(`should render correctly`, () => {
    const { container } = render(testingButton);
    expect(container).toMatchSnapshot();
  });

  it(`should have a functioning onClick handler`, () => {
    const { container } = render(testingButton);
    const button = container.querySelector(`button`);
    act(() => {
      if (!button) {
        throw new Error(`No button found`);
      }
      fireEvent.click(button);
    });
    expect(callback).toHaveBeenCalled();
  });

  it(`should be able to animate correctly`, () => {
    jest.useFakeTimers();
    jest.spyOn(global, `setTimeout`);
    const { container } = render(testingButton);
    const button = container.querySelector(`button`);
    act(() => {
      if (!button) {
        throw new Error(`No button found`);
      }
      fireEvent.click(button);
    });
    expect(container.querySelector(`.animBarActive`)).toBeTruthy();
    act(() => {
      jest.runOnlyPendingTimers();
    });
    expect(container.querySelector(`.animBarActive`)).toBeFalsy();
    expect(setTimeout).toHaveBeenCalledTimes(1);
  });

  it(`should join class names correctly`, () => {
    const { container } = render(testingButtonWithClassName);
    expect(
      container.querySelector(`button`)?.classList[0] === `test-class`,
    ).toBeTruthy();
  });
});
