export enum EventNames {
  copyTileMarkers = `copy_tile_markers`,
}

type Event = {
  action: EventNames;
  params: {
    [key: string]: string | number;
  };
};

export const event = (event: Event) => {
  if (typeof window.gtag === `function`) {
    window.gtag(`event`, event.action, event.params);
  }
};
