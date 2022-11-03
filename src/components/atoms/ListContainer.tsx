import styles from '@/components/atoms/ListContainer.module.css';

export default function ListContainer({
  children,
  customRef,
}: {
  children: React.ReactNode;
  customRef?: React.RefObject<HTMLDivElement>;
}) {
  return (
    <div ref={customRef} className={styles.innerListContainer}>
      {children}
    </div>
  );
}
