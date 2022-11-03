import Button from '@/components/atoms/Button';
import Image from 'next/image';
import styles from '@/components/molecules/TileEntityCardCompact.module.css';
import copy from 'copy-to-clipboard';
import Link from 'next/link';
import { event, EventNames } from '@/api/gtag';
import { useState } from 'react';
import { TileEntityCardCompactProps } from '@/types';

export default function TileEntityCardCompact({
  entity,
}: TileEntityCardCompactProps) {
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
          width={120}
          height={120}
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
        </div>
        <div className={styles.tileInteraction}>
          <p className={styles.tileCount}>{entity.tiles.length} tile markers</p>
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
        </div>
      </div>
    </div>
  );
}
