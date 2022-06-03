import styles from '@/components/atoms/StripContainer.module.css';

export default function StripContainer({
  children,
}: {
  children: React.ReactNode;
}) {
  return <div className={styles.innerContainer}>{children}</div>;
}
