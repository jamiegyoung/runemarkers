import styles from '@/components/atoms/Title.module.css';
import Link from './Link';

export default function Title() {
  return (
    <Link className={styles.titleLink} href="/">
      <header className={styles.title}>
        R<span>une</span>
        <span className={styles.titleAlt}>
          M<span>arkers</span>
        </span>
        <span>.net</span>
      </header>
    </Link>
  );
}
