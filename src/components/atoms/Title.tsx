import styles from '@/components/atoms/Title.module.css';

type TitleProps = {
  size?: number;
};

export default function Title({ size }: TitleProps) {
  return (
    <a
      href="./"
      style={{
        background: `transparent`,
        border: `none`,
        color: `#f5f3f5`,
        textDecoration: `none`,
        display: `flex`,
      }}
    >
      <header
        style={{
          fontSize: size ? size : `2rem`,
        }}
        className={styles.title}
      >
        R<span>une</span>
        <span className={styles.titleAlt}>
          M<span>arkers</span>
        </span>
      </header>
    </a>
  );
}
