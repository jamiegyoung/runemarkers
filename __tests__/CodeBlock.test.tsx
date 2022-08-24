import '@testing-library/jest-dom';
import { act, fireEvent, render } from '@testing-library/react';
import CodeBlock from '@/components/atoms/CodeBlock';
import { Tile } from '@/types';

const testTiles: Tile[] = [
  {
    regionId: 7322,
    regionX: 47,
    regionY: 43,
    z: 0,
    color: `#FFFAFF00`,
  },
  {
    regionId: 7322,
    regionX: 47,
    regionY: 48,
    z: 0,
    color: `#FFFAFF00`,
  },
  {
    regionId: 7322,
    regionX: 52,
    regionY: 48,
    z: 0,
    color: `#FFFAFF00`,
  },
  {
    regionId: 7322,
    regionX: 52,
    regionY: 43,
    z: 0,
    color: `#FFFAFF00`,
  },
];

// const otherTiles: Tile[] = [
//   {
//     regionId: 5536,
//     regionX: 27,
//     regionY: 38,
//     z: 0,
//     color: `#FFFFFF00`,
//   },
// ];

const testCodeBlock = (tiles: Tile[], truncateLength?: number) => (
  <CodeBlock truncateLength={truncateLength} tiles={tiles} />
);

describe(`CodeBlock`, () => {
  it(`should render correctly`, () => {
    const { container } = render(testCodeBlock(testTiles, 5));
    expect(container).toMatchSnapshot();
  });

  it(`should render correctly with pretty`, () => {
    const { container } = render(testCodeBlock(testTiles, 5));
    const prettyButton = container.querySelector(`.togglePretty`);
    if (!prettyButton) {
      throw new Error(`prettyButton not found`);
    }
    const code = container.querySelector(`.code`);
    if (!code) {
      throw new Error(`code not found`);
    }

    expect(code.textContent).not.toContain(`\t`);

    act(() => {
      fireEvent.click(prettyButton);
    });

    // only pretty code should contain tabs
    expect(code.textContent).toContain(`\t`);
  });

  // it(`should change when the tiles change`, () => {
  //   const { container, rerender } = render(testCodeBlock(testTiles, 5));
  //   expect(container).toMatchSnapshot();
  //   rerender(testCodeBlock(otherTiles, 5));
  //   expect(container).toMatchSnapshot();
  // });

  it(`should truncate lines when pretty correctly`, () => {
    const { container } = render(testCodeBlock(testTiles));
    const code = container.querySelector(`.code`);
    if (!code) {
      throw new Error(`Code block not found`);
    }
    if (!code.textContent) {
      throw new Error(`Code block text content not found`);
    }
    const prettyButton = container.querySelector(`.togglePretty`);
    if (!prettyButton) {
      throw new Error(`prettyButton not found`);
    }
    act(() => {
      fireEvent.click(prettyButton);
    });
    // one extra line is appended to the end
    expect(code.textContent.split(`\n`)).toHaveLength(11);
    expect(code.textContent).toContain(`...`);
  });

  it(`should not have the toggle if the text is short enough`, () => {
    const { container } = render(testCodeBlock(testTiles, 500));
    expect(container.querySelector(`.toggleShowAll`)).toBeFalsy();
  });

  it(`should show all on click`, () => {
    const { container } = render(testCodeBlock(testTiles, 5));
    const buttons = container.querySelectorAll(`.toggleShowAll`);
    if (!buttons) {
      throw new Error(`No button found`);
    }
    act(() => {
      fireEvent.click(buttons[0]);
    });
    expect(buttons[0].textContent).toBe(`show less`);
    const newButtons = container.querySelectorAll(`.toggleShowAll`);
    expect(newButtons.length).toBe(2);
    act(() => {
      fireEvent.click(newButtons[1]);
    });
    expect(newButtons[1]).not.toBeInTheDocument();
    expect(newButtons[0].textContent).toContain(`show all`);
  });

  it(`should become pretty when the option is clicked`, () => {
    const { container } = render(testCodeBlock(testTiles, 5));
    const code = container.querySelector(`.code`);
    // check it isn't already pretty
    expect(code).not.toContain(`\t`);

    const button = container.querySelector(`.togglePretty`);
    if (!button) {
      throw new Error(`No button found`);
    }
    act(() => {
      fireEvent.click(button);
    });
    if (!code) {
      throw new Error(`No code found`);
    }
    expect(code.textContent).toContain(`\t`);
  });
});
