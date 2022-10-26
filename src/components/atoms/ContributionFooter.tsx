import Image from 'next/image';
import styles from '@/components/atoms/ContributionFooter.module.css';

export default function ContributionFooter() {
  return (
    <footer className={styles.footer}>
      <div className={styles.footerItem}>
        <p>
          Made by
          <a
            href="https://discord.com/users/167850724976361472"
            className={styles.link}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Image src="/discord.png" alt="GitHub" width={21} height={24} />
            <span>Jam#0001</span>
          </a>
        </p>
      </div>
      <div className={styles.footerItem}>
        <a
          href="https://github.com/jamiegyoung/runemarkers"
          className={styles.link}
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image src="/github.png" alt="GitHub" width={20} height={20} />
          <span>contribute code / tiles</span>
        </a>
      </div>
    </footer>
  );
}
