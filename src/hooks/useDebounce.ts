import { useRef, useEffect } from 'react';

export default function useDebounce<T extends (...args: any[]) => void>(
  fn: T,
  delay: number,
): (...args: Parameters<T>) => void {
  const timeoutId = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    return () => {
      if (timeoutId.current !== null) clearTimeout(timeoutId.current);
    };
  }, []);

  return (...args: Parameters<T>) => {
    if (timeoutId.current !== null) clearTimeout(timeoutId.current);
    timeoutId.current = setTimeout(() => fn(...args), delay);
  };
}
