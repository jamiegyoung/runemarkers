import Image from 'next/image';
import styles from '@/components/atoms/GitHubFooter.module.css';

export default function ContributionFooter() {
  return (
    <footer className={styles.footer}>
      <a href="https://github.com/jamiegyoung/runemarkers">
        <Image src="/github.png" alt="GitHub" width={20} height={20} />
        <span style={{ marginLeft: 5 }}>contribute code / tiles</span>
      </a>
    </footer>
  );
}
