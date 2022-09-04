import Search from '../molecules/Search';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import { TileEntity } from '@/types';
import styles from '@/components/organisms/TileEntityList.module.css';
import StripContainer from '../atoms/StripContainer';
import { useContext, useEffect, useRef } from 'react';
import { SearchContext } from '@/pages';

type TileEntityListProps = {
  list: TileEntity[];
};

export default function TileEntityList({ list }: TileEntityListProps) {
  const ref = useRef<HTMLDivElement>(null);
  const [searchVal] = useContext(SearchContext);

  useEffect(() => {
    if (ref.current) {
      ref.current.scrollTop = 0;
    }
  }, [searchVal]);

  return (
    <div className={styles.container}>
      <p
        style={{
          padding: `0`,
          color: `#C7C7C7`,
          textAlign: `center`,
          minWidth: 280,
          marginBottom: 15,
          maxWidth: 600,
          width: `80%`,
        }}
      >
        Make sure the &quot;Ground Markers&quot; plugin is enabled on RuneLite
        with &quot;Remember color per tile&quot; enabled. Copy the tile markers
        to your clipboard, right-click the world map in-game and select
        &quot;import&quot;.
      </p>
      <Search />
      <StripContainer customRef={ref}>
        {list.length > 0 ? (
          list.map((i) => <TileEntityCard key={i.safeURI} entity={i} />)
        ) : (
          <p style={{ textAlign: `center` }}>No results found</p>
        )}
      </StripContainer>
    </div>
  );
}
