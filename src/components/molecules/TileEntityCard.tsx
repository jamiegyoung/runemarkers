import Button from '@/components/atoms/Button';
import InfoLink from '@/components/atoms/InfoLink';
import { TileEntity } from '@/types';
import Image from 'next/image';
import styles from '@/components/molecules/TileEntityCard.module.css';
import copy from 'copy-to-clipboard';
import Link from 'next/link';
import { event, EventNames } from '@/api/gtag';

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
      <Link
        href={`/${encodeURIComponent(entity.safeURI)}`}
        className={styles.pictureLink}
      >
        <a className={styles.imageContainer}>
          <Image
            className={styles.image}
            width={140}
            height={140}
            src={entity.thumbnail}
            alt={`${entity.name}.png`}
          />
        </a>
      </Link>
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <Link href={`/${encodeURIComponent(entity.safeURI)}`}>
            <a className={styles.name}>
              {entity.name}
              {entity.subcategory ? (
                <>
                  {` `}
                  <span style={{ whiteSpace: `nowrap` }}>
                    ({entity.subcategory})
                  </span>
                </>
              ) : null}
            </a>
          </Link>
          {entity.altName ? (
            <Link href={`/${encodeURIComponent(entity.safeURI)}`}>
              <a className={styles.altName}>
                {entity.altName}
                {entity.subcategory ? ` (${entity.subcategory}) ` : null}
              </a>
            </Link>
          ) : null}
        </div>
        <div>
          <p className={styles.tileCount}>{entity.tiles.length} tile markers</p>
          <div className={styles.tileInteraction}>
            <Button
              className={styles.button}
              onClick={() => {
                copy(JSON.stringify(entity.tiles));
                event({
                  action: EventNames.copyTileMarkers,
                  params: {
                    event_category: `engagement`,
                    event_label: entity.safeURI,
                    tiles_copied: entity.safeURI,
                  },
                });
              }}
            >
              Copy
            </Button>
            {hideInfoButton ? null : (
              <InfoLink href={`/${encodeURIComponent(entity.safeURI)}`} />
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
