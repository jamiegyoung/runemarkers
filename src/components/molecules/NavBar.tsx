import Title from '@/components/atoms/Title';
import styles from '@/components/molecules/NavBar.module.css';

export default function NavBar() {
  return (
    <div className={styles.navbar}>
      <Title />
    </div>
  );
}
