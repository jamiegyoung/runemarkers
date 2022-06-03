import Title from "../atoms/Title";
import styles from './NavBar.module.css';
import Search from './Search';

export default function NavBar() {
  return (
    <div className={styles.navbar}>
      <Title short/>
      <Search />
    </div>
  );
}
