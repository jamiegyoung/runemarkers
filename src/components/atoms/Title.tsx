import styles from '@/components/atoms/Title.module.css';
import useWindowDimensions from '@/hooks/useWindowDimensions';
import { useEffect, useState } from 'react';

type TitleProps = {
  size?: number;
};

export default function Title({ size }: TitleProps) {
  const [short, setShort] = useState(false);
  const { width } = useWindowDimensions();
  useEffect(() => {
    if (width < 400) {
      setShort(true);
      return;
    }
    setShort(false);
  }, [setShort, width]);

  return (
    <a
      href="./"
      style={{
        background: `transparent`,
        border: `none`,
        color: `#f5f3f5`,
        textDecoration: `none`,
      }}
    >
      <header
        style={{
          fontSize: size ? size : `2rem`,
        }}
        className={styles.title}
      >
        {short ? `R` : `Rune`}
        <span>{short ? `M` : `Markers`}</span>
      </header>
    </a>
  );
}
