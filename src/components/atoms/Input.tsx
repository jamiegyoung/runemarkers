import { SearchContext } from '@/pages';
import styles from '@/components/atoms/Input.module.css';
import { useContext, useState } from 'react';

export default function Input() {
  const [searching, setSearching] = useState(false);
  const [searchVal, setSearchVal] = useContext(SearchContext);

  return (
    <input
      type="text"
      placeholder={searching ? `` : `type here to search`}
      className={styles.input}
      onFocus={() => setSearching(true)}
      onBlur={() => {
        if (searchVal === ``) {
          setSearching(false);
        }
      }}
      style={{
        width: searching ? 250 : 175,
        minWidth: 50,
        maxWidth: 500,
      }}
      onChange={(e) => {
        if (setSearchVal) {
          setSearchVal(e.target.value);
        }
      }}
    />
  );
}
