import { useContext } from 'react';
import { SearchContext } from '../App';
import Input from "../atoms/Input";
import SearchIcon from "../atoms/SearchIcon";

export default function Search() {
  const [searchVal, setSearchVal] = useContext(SearchContext);;

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
      }}
    >
      <Input />
      <SearchIcon/>
    </div>
  );
}
