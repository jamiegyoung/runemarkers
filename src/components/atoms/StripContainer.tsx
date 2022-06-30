import styles from '@/components/atoms/StripContainer.module.css';

export default function StripContainer({
  children,
  customRef,
}: {
  children: React.ReactNode;
  customRef: React.RefObject<HTMLDivElement>;
}) {
  return (
    <div ref={customRef} className={styles.innerContainer}>
      {children}
    </div>
  );
}
