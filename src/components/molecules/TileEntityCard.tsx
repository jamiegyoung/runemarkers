import Button from '@/components/atoms/Button';
import InfoLink from '@/components/atoms/InfoLink';
import { TileEntity } from '@/types';
import Image from 'next/image';
import styles from '@/components/molecules/TileEntityCard.module.css';
import copy from 'copy-to-clipboard';
import Link from 'next/link';
import { event, EventNames } from '@/api/gtag';
import { useState } from 'react';

type TileEntityCardProps = {
  entity: TileEntity;
  hideInfoButton?: boolean;
};

export default function TileEntityCard({
  entity,
  hideInfoButton,
}: TileEntityCardProps) {
  const [imageHidden, setImageHidden] = useState(true);

  return (
    <div className={styles.card}>
      <Link
        className={styles.imageContainer}
        href={`/${encodeURIComponent(entity.safeURI)}`}
      >
        <Image
          onLoad={() => setImageHidden(false)}
          className={styles.image}
          style={{ opacity: imageHidden ? 0 : 1 }}
          width={100}
          height={100}
          src={entity.thumbnail}
          alt={`${entity.name} image`}
        />
      </Link>
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <Link
            className={styles.name}
            href={`/${encodeURIComponent(entity.safeURI)}`}
          >
            {entity.name}
            {entity.subcategory ? (
              <>
                {` `}
                <span style={{ whiteSpace: `nowrap` }}>
                  ({entity.subcategory})
                </span>
              </>
            ) : null}
          </Link>
          {entity.altName ? (
            <Link
              className={styles.altName}
              href={`/${encodeURIComponent(entity.safeURI)}`}
            >
              {entity.altName}
              {entity.subcategory ? ` (${entity.subcategory}) ` : null}
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
              <InfoLink
                href={`/${encodeURIComponent(entity.safeURI)}`}
                name={entity.name}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
