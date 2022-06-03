import Search from '../molecules/Search';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import { TileEntity } from '@/types';
import styles from '@/components/organisms/TileEntityList.module.css';
import StripContainer from '../atoms/StripContainer';

type TileEntityListProps = {
  list: TileEntity[];
};

export default function TileEntityList({ list }: TileEntityListProps) {
  return (
    <div className={styles.container}>
      <p
        style={{
          padding: `0 20px`,
          color: `#C7C7C7`,
          maxWidth: `500px`,
          textAlign: `center`,
        }}
      >
        Make sure the “Gound Markers” plugin is enabled in runelite, copy the
        tile markers to your clipboard, right click the world map on runelite
        and then select “import”
      </p>
      <Search />
      <StripContainer>
        {list.map((i) => (
          <TileEntityCard key={i.name} entity={i} />
        ))}
      </StripContainer>
    </div>
  );
}
