import styles from './Button.module.css';

export default function Button(
  buttonProps: React.ButtonHTMLAttributes<HTMLButtonElement>
) {
  return <button className={styles.button} {...buttonProps}>{buttonProps.children}</button>;
}
