import { Source } from '@/types';
import styles from '@/components/molecules/TilesSource.module.css';
import Modified from '@/components/atoms/Modified';
import Link from '../atoms/Link';

export default function TilesSource({ source }: { source: Source }) {
  return (
    <p className={styles.tilesSourceParagraph}>
      <Link className={styles.link} href={source.link} newTab>
        {source.name}
      </Link>
      {source.modified ? <Modified>{source.modified}</Modified> : null}
    </p>
  );
}
