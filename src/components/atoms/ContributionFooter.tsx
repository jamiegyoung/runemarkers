import Image from 'next/image';
import styles from '@/components/atoms/ContributionFooter.module.css';

export default function ContributionFooter() {
  return (
    <footer className={styles.footer}>
      <p>Made by</p>
      <a href="https://discord.com/users/167850724976361472">
        <Image src="/discord.png" alt="GitHub" width={20} height={20} />
        <span>Jam#0001</span>
      </a>
      <a href="https://github.com/jamiegyoung/runemarkers">
        <Image src="/github.png" alt="GitHub" width={20} height={20} />
        <span>contribute code / tiles</span>
      </a>
    </footer>
  );
}
