import { Source } from '@/types';
import styles from '@/components/molecules/TilesSource.module.css';
import Modified from '@/components/atoms/Modified';

export default function TilesSource({ source }: { source: Source }) {
  return (
    <div className={styles.tilesSourceParagraph}>
      <a
        className={styles.link}
        href={source.link}
        target="_blank"
        rel="noopener noreferrer"
      >
        {source.name}
      </a>
      {source.modified ? <Modified>{source.modified}</Modified> : null}
    </div>
  );
}
