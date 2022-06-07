import Fuse from 'fuse.js';
import { TileEntity } from '@/types';
import { createContext, useEffect, useState } from 'react';
import TileEntityList from '@/components/organisms/TileEntityList';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/Home.module.css';
import { tileData } from '@/api/tileJson';

const fuseOptions = {
  // the keys of the objects to search
  keys: [
    { name: `name`, weight: 0.7 },
    { name: `tags`, weight: 0.3 },
  ],
  threshold: 0.4,
};

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
      <main className={styles.main}>
        <SearchContext.Provider value={[searchVal, setSearchVal]}>
          <NavBar />
          <TileEntityList list={searchRes} />
        </SearchContext.Provider>
      </main>
    </div>
  );
}
