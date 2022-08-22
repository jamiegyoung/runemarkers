import { EventNames } from '@/api/gtag';

export {};

type GtagFn = (
  command: string,
  action: EventNames,
  params: EventParams,
) => void;

declare global {
  interface Window {
    gtag: undefined | GtagFn;
  }
}
