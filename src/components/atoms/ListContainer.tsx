import styles from '@/components/atoms/ListContainer.module.css';
import { ReactNode, RefObject } from 'react';

export default function ListContainer({
  children,
  customRef,
}: {
  children: ReactNode;
  customRef?: RefObject<HTMLDivElement>;
}) {
  return (
    <div ref={customRef} className={styles.innerListContainer}>
      {children}
    </div>
  );
}
