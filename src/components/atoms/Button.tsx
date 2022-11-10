import { ButtonHTMLAttributes, useEffect, useState } from 'react';
import styles from '@/components/atoms/Button.module.css';

export default function Button(
  buttonProps: ButtonHTMLAttributes<HTMLButtonElement>,
) {
  const [isAnimating, setIsAnimating] = useState(false);
  useEffect(() => {
    if (!isAnimating) {
      return;
    }
    const timeout = setTimeout(() => {
      setIsAnimating(false);
    }, 1100);
    return () => clearTimeout(timeout);
  }, [isAnimating]);

  return (
    <button
      {...buttonProps}
      className={
        buttonProps.className
          ? [buttonProps.className, styles.button].join(` `)
          : styles.button
      }
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
          isAnimating ? styles.animBarActive : ``,
        ].join(` `)}
      ></div>
    </button>
  );
}
