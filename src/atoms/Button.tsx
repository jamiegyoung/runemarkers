import { useEffect, useState } from "react";
import styles from "./Button.module.css";

export default function Button(
  buttonProps: React.ButtonHTMLAttributes<HTMLButtonElement>
) {
  const [isAnimating, setIsAnimating] = useState(false);
  useEffect(() => {
    const timeout = setTimeout(() => {
      setIsAnimating(false);
    }, 1100);
    return () => clearTimeout(timeout);
  }, [isAnimating]);

  return (
    <button
      className={styles.button}
      {...buttonProps}
      onClick={(e) => {
        if (buttonProps.onClick) {
          buttonProps.onClick(e);
        }
        setIsAnimating(true);
      }}
    >
      {buttonProps.children}
      <div
        className={[
          styles.animBar,
          isAnimating ? styles.animBarActive : "",
        ].join(" ")}
      ></div>
    </button>
  );
}
