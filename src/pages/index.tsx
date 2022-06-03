import Head from 'next/head';
import Fuse from 'fuse.js';
import { TileEntity } from '@/types';
import tileJson from '../tiles.json';
import { createContext, useEffect, useState } from 'react';
import TileEntityList from '@/components/organisms/TileEntityList';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/Home.module.css';
// import TileEntityList from "./organisms/TileEntityList";

const fuseOptions = {
  // debugging
  // includeScore: true,
  // the keys of the objects to search
  keys: [
    { name: `name`, weight: 0.7 },
    { name: `tags`, weight: 0.3 },
  ],
  threshold: 0.4,
};

// Assert the structure of the tileJson
const tileData: TileEntity[] = tileJson;

const fuse = new Fuse(tileData, fuseOptions);

export const SearchContext = createContext<
  [string | undefined, React.Dispatch<React.SetStateAction<string>> | undefined]
>([undefined, undefined]);

export default function Home() {
  const [searchVal, setSearchVal] = useState(``);
  const [searchRes, setSearchRes] = useState<TileEntity[]>(tileData);

  useEffect(() => {
    if (searchVal.length > 0) {
      setSearchRes(fuse.search(searchVal).map((res) => res.item));
      return;
    }
    setSearchRes(tileData);
  }, [searchVal]);
  return (
    <div>
      <Head>
        <title>RuneMarkers - A Collection of RuneLite Ground Markers</title>
        <meta
          name="description"
          content="Find and import RuneLite ground markers for oldschool runescape. Easily search for ground markers by boss, activity or location."
        />
        <meta name="author" content="RuneMarkers" />
        <meta property="og:site_name" content="RuneMarkers" />
        <meta
          property="og:title"
          content="RuneMarkers - A Collection of RuneLite Ground Markers"
        />
        <meta
          property="og:description"
          content="Find and import RuneLite ground markers for oldschool runescape."
        />
        <meta property="og:type" content="website" />
        <meta property="og:url" content="https://runemarkers.net" />
        <meta property="og:image" content="/logo1024.png" />
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="theme-color" content="#262624" />
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/apple-touch-icon.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon-16x16.png"
        />
        <link rel="manifest" href="/site.webmanifest" />
      </Head>

      <main className={styles.main}>
        <SearchContext.Provider value={[searchVal, setSearchVal]}>
          <NavBar />
          <TileEntityList list={searchRes} />
        </SearchContext.Provider>
      </main>
    </div>
  );
}
