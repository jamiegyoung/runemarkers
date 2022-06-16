import { NextSeoProps } from 'next-seo';

const nextSeoOptions: NextSeoProps = {
  defaultTitle: `RuneMarkers - A Collection of RuneLite Ground Markers`,
  description: `A collection of ground (tile) markers for RuneLite. Find and import ground (tile) markers for different Oldschool RuneScape activities.`,
  titleTemplate: `RuneMarkers - %s`,
  canonical: `https://runemarkers.net`,
  openGraph: {
    type: `website`,
    locale: `en_US`,
    url: `https://runemarkers.net`,
    title: `RuneMarkers - A Collection of RuneLite Ground Markers`,
    description: `A collection of ground markers for RuneLite. Find and import ground markers for different Oldschool RuneScape activities.`,
    site_name: `RuneMarkers`,
  },
  twitter: {
    handle: `@jamgyo`,
    site: `@runemarkers`,
    cardType: `summary_large_image`,
  },
};

export const defaultImages = [
  {
    url: `/logo1024-background.png`,
    width: 1024,
    height: 1024,
    alt: `RuneMarkers - A Collection of RuneLite Ground Markers`,
  },
  {
    url: `/logo512-background.png`,
    width: 512,
    height: 512,
    alt: `RuneMarkers - A Collection of RuneLite Ground Markers`,
  },
  {
    url: `/android-chrome-512x512.png`,
    width: 512,
    height: 512,
    alt: `RuneMarkers - A Collection of RuneLite Ground Markers`,
  },
  {
    url: `/android-chrome-192x192.png`,
    width: 192,
    height: 192,
    alt: `RuneMarkers - A Collection of RuneLite Ground Markers`,
  },
];

export default nextSeoOptions;
