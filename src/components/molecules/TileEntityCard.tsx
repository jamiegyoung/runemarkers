import Button from '@/components/atoms/Button';
import InfoButton from '@/components/atoms/InfoLink';
import { TileEntity } from '@/types';
import Image from 'next/image';
import styles from '@/components/molecules/TileEntityCard.module.css';
import copy from 'copy-to-clipboard';

type TileEntityCardProps = {
  entity: TileEntity;
  hideInfoButton?: boolean;
};

export default function TileEntityCard({
  entity,
  hideInfoButton,
}: TileEntityCardProps) {
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
          width={140}
          height={140}
          src={entity.thumbnail}
          alt={`${entity.name}.png`}
        />
      </a>
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <h1>{entity.name}</h1>
          {entity.altName ? (
            <h2 className={styles.altName}>{entity.altName}</h2>
          ) : null}
        </div>
        <div>
          <p className={styles.tileCount}>{entity.tiles.length} tile markers</p>
          <div className={styles.tileInteraction}>
            <Button
              className={styles.button}
              onClick={() => {
                copy(JSON.stringify(entity.tiles));
              }}
            >
              Copy
            </Button>
            {hideInfoButton ? null : (
              <InfoButton href={`/${encodeURIComponent(entity.safeURI)}`} />
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
