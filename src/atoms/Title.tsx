import styles from "./Title.module.css";

type TitleProps = {
  short?: boolean;
  size?: number;
};

export default function Title({ short, size }: TitleProps) {
  return (
    <header
      style={{
        fontSize: size ? size : "2rem",
      }}
      className={styles.title}
    >
      {short ? "R" : "Rune"}
      <span>{short ? "M" : "Markers"}</span>
    </header>
  );
}
