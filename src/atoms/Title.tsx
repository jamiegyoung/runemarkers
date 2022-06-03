import styles from "./Title.module.css";

type TitleProps = {
  short?: boolean;
  size?: number;
};

export default function Title({ short, size }: TitleProps) {
  return (
    <a href='./' style={{
      background: "transparent",
      border: "none",
      color: "#f5f3f5",
      textDecoration: "none",
    }}>
      <header
        style={{
          fontSize: size ? size : "2rem",
        }}
        className={styles.title}
      >
        {short ? "R" : "Rune"}
        <span>{short ? "M" : "Markers"}</span>
      </header>
    </a>
  );
}
