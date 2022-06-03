import Title from "../atoms/Title";
import useWindowDimensions from "../hooks/useWindowDimensions";
import styles from "./NavBar.module.css";

export default function NavBar() {
  const { width } = useWindowDimensions();
  return (
    <div className={styles.navbar}>
      <Title short={width < 600} />
    </div>
  );
}
