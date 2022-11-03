import useWindowDimensions from '@/hooks/useWindowDimensions';
import '@testing-library/jest-dom';
import { act, renderHook } from '@testing-library/react';

describe(`useWindowDimensions`, () => {
  it(`should return the correct window dimensions`, () => {
    const { result } = renderHook(() => useWindowDimensions());
    expect(result.current.width).toBe(1024);
    expect(result.current.height).toBe(768);
  });

  it(`should return the correct window dimensions when the window is resized`, () => {
    const { result } = renderHook(() => useWindowDimensions());
    expect(result.current.width).toBe(1024);
    expect(result.current.height).toBe(768);
    Object.defineProperty(window, `innerWidth`, { writable: true, value: 800 });
    Object.defineProperty(window, `innerHeight`, {
      writable: true,
      value: 600,
    });
    act(() => {
      window.dispatchEvent(new Event(`resize`));
    });
    expect(result.current.width).toBe(800);
    expect(result.current.height).toBe(600);
  });
});
