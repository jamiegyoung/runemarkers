import Button from '@/components/atoms/Button';
import InfoButton from '@/components/atoms/InfoButton';
import { TileEntity } from '@/types';
import Image from 'next/image';
import styles from '@/components/molecules/TileEntityCard.module.css';
import copy from 'copy-to-clipboard';

type TileEntityCardProps = {
  entity: TileEntity;
};

export default function TileEntityCard({ entity }: TileEntityCardProps) {
  return (
    <div className={styles.card}>
      <a
        href={entity.wiki}
        target="_blank"
        rel="noopener noreferrer"
        className={styles.pictureLink}
      >
        <Image
          className={styles.image}
          width={200}
          height={200}
          src={entity.thumbnail}
          alt={`${entity.name}.png`}
        />
      </a>
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <h1>{entity.name}</h1>
          <p>{entity.tiles.length} tile markers</p>
        </div>
        <div className={styles.tileInteraction}>
          <Button
            className={styles.button}
            onClick={() => {
              copy(JSON.stringify(entity.tiles));
            }}
          >
            Copy
          </Button>
          <InfoButton />
        </div>
      </div>
    </div>
  );
}
