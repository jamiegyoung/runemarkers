import styles from '@/components/atoms/GridContainer.module.css';
import { ReactNode, RefObject } from 'react';

export default function GridContainer({
  children,
  customRef,
}: {
  children: ReactNode;
  customRef?: RefObject<HTMLDivElement>;
}) {
  return (
    <div ref={customRef} className={styles.innerGridContainer}>
      {children}
    </div>
  );
}
