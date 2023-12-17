import styles from '@/components/atoms/Modified.module.css';

export default function Modified({ children }: { children: string }) {
  return (
    <span className={styles.tooltip}>
      (modified)
      <span className={styles.tooltipWrapper}>
        <span className={styles.tooltipText}>{children}</span>
      </span>
    </span>
  );
}
