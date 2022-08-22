export enum EventNames {
  copyTileMarkers = `copy_tile_markers`,
}

export type EventParams = {
  [key: string]: string | number | undefined;
  event_category?: string;
  event_label?: string;
  value?: number;
};

type Event = {
  action: EventNames;
  params: EventParams;
};

export const event = (event: Event) => {
  if (typeof window.gtag === `function`) {
    window.gtag(`event`, event.action, event.params);
  }
};
