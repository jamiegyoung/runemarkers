import { SearchContext } from '@/pages';
import styles from '@/components/atoms/Input.module.css';
import { useContext, useEffect, useRef, useState } from 'react';
import useDebounce from '@/hooks/useDebounce';

export default function Input() {
  const [searching, setSearching] = useState(false);
  const [searchVal, setSearchVal] = useContext(SearchContext);
  const inputRef = useRef<HTMLInputElement>(null);
  const [inputWidth, setInputWidth] = useState(0);

  const updateWidth = () => {
    if (!inputRef.current) return;
    setInputWidth(inputRef.current.offsetWidth);
  };

  const updateWidthDebounced = useDebounce(updateWidth, 100);

  useEffect(() => {
    updateWidth();
    window.addEventListener(`resize`, updateWidthDebounced);
  });

  return (
    <>
      <input
        ref={inputRef}
        type="text"
        placeholder={
          searching
            ? ``
            : inputWidth > 230
              ? `type here to search`
              : `search tiles`
        }
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
