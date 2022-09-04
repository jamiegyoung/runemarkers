import { SearchContext } from '@/pages';
import styles from '@/components/atoms/Input.module.css';
import { useContext, useState } from 'react';

export default function Input() {
  const [searching, setSearching] = useState(false);
  const [searchVal, setSearchVal] = useContext(SearchContext);

  return (
    <>
      <input
        type="text"
        placeholder={searching ? `` : `type here to search`}
        className={[styles.input, searching ? styles.active : null].join(` `)}
        value={searchVal}
        onFocus={() => setSearching(true)}
        onBlur={() => {
          if (searchVal === ``) {
            setSearching(false);
          }
        }}
        onChange={(e) => {
          if (setSearchVal) {
            setSearchVal(e.target.value);
          }
        }}
      />
      <button
        className={[styles.clearButton, searching ? styles.active : null].join(
          ` `,
        )}
        onClick={() => {
          if (setSearchVal) {
            setSearching(false);
            setSearchVal(``);
          }
        }}
      >
        <svg xmlns="http://www.w3.org/2000/svg" height="24" width="24">
          <path d="M6.4 19.45 4.55 17.6l5.6-5.6-5.6-5.6L6.4 4.55l5.6 5.6 5.6-5.6 1.85 1.85-5.6 5.6 5.6 5.6-1.85 1.85-5.6-5.6Z" />
        </svg>
      </button>
    </>
  );
}
