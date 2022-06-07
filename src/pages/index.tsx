import Fuse from 'fuse.js';
import { TileEntity } from '@/types';
import { createContext, useEffect, useState } from 'react';
import TileEntityList from '@/components/organisms/TileEntityList';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/Home.module.css';
import { getTileData } from '@/api/tiles';

const fuseOptions = {
  // the keys of the objects to search
  keys: [
    { name: `name`, weight: 0.7 },
    { name: `tags`, weight: 0.3 },
  ],
  threshold: 0.4,
};

export async function getStaticProps() {
  const data = getTileData();
  return {
    props: { tileData: data },
  };
}

export const SearchContext = createContext<
  [string | undefined, React.Dispatch<React.SetStateAction<string>> | undefined]
>([undefined, undefined]);

export default function Home({ tileData }: { tileData: TileEntity[] }) {
  const [searchVal, setSearchVal] = useState(``);
  const [searchRes, setSearchRes] = useState<TileEntity[]>(tileData);

  useEffect(() => {
    const fuse = new Fuse(tileData, fuseOptions);
    if (searchVal.length > 0) {
      setSearchRes(fuse.search(searchVal).map((res) => res.item));
      return;
    }
    setSearchRes(tileData);
  }, [searchVal, tileData]);
  return (
    <div>
      <main className={styles.main}>
        <SearchContext.Provider value={[searchVal, setSearchVal]}>
          <NavBar />
          <TileEntityList list={searchRes} />
        </SearchContext.Provider>
      </main>
    </div>
  );
}
