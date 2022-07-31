import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import CodeBlock from '@/components/atoms/CodeBlock';

const testText = `123456789`;

const testCodeBlock = (truncateLength: number | undefined) => (
  <CodeBlock truncateLength={truncateLength} text={testText} />
);

describe(`CodeBlock`, () => {
  it(`should render correctly`, () => {
    const { container } = render(testCodeBlock(5));
    expect(container).toMatchSnapshot();
  });

  it(`should not have the toggle if the text is short enough`, () => {
    const { container } = render(testCodeBlock(undefined));
    expect(container.querySelector(`#toggleShowAll`)).toBeFalsy();
  });

  it(`should show all on click`, () => {
    const { container } = render(testCodeBlock(5));
    const button = container.querySelector(`#toggleShowAll`);
    if (!button) {
      throw new Error(`No button found`);
    }
    act(() => {
      fireEvent.click(button);
    });
    expect(button.textContent).toBe(`Show Less`);
  });
});
