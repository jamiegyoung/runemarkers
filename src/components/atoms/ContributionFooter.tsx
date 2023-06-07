import Image from 'next/image';
import styles from '@/components/atoms/ContributionFooter.module.css';
import Link from './Link';

export default function ContributionFooter() {
  return (
    <footer className={styles.footer}>
      <div className={styles.footerItem}>
        <p>
          Made by
          <Link href="https://discord.com/users/167850724976361472" newTab>
            <Image src="/discord.png" alt="GitHub" width={20} height={17} />
            <span className={styles.linkText}>jamgyo</span>
          </Link>
        </p>
      </div>
      <div className={styles.footerItem}>
        <Link href="https://github.com/jamiegyoung/runemarkers" newTab>
          <Image src="/github.png" alt="GitHub" width={20} height={20} />
          <span className={styles.linkText}>contribute code / tiles</span>
        </Link>
      </div>
    </footer>
  );
}
