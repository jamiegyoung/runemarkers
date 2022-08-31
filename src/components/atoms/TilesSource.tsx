import { Source } from '@/types';
import styles from '@/components/atoms/TilesSource.module.css';
import Modified from './Modified';

export default function TilesSource({ source }: { source: Source }) {
  return (
    <div className={styles.tilesSourceContainer}>
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
