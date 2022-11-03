import styles from '@/components/atoms/GridContainer.module.css';

export default function GridContainer({
  children,
  customRef,
}: {
  children: React.ReactNode;
  customRef?: React.RefObject<HTMLDivElement>;
}) {
  return (
    <div ref={customRef} className={styles.innerGridContainer}>
      {children}
    </div>
  );
}
