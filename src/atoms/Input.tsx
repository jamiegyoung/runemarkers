import { useContext, useState } from "react";
import { SearchContext } from "../App";
import styles from "./Input.module.css";

export default function Input() {
  const [searching, setSearching] = useState(false);
  const [searchVal, setSearchVal] = useContext(SearchContext);;

  return (
    <input
      type="text"
      placeholder="Search"
      className={styles.input}
      onFocus={() => setSearching(true)}
      onBlur={() => {
        if (searchVal === "") {
          setSearching(false);
        }
      }}
      style={{
        width: searching ? "7rem" : "3rem",
      }}
      onChange={(e) => {
        if (setSearchVal) {
          setSearchVal(e.target.value);
        }
      }}
    />
  );
}
