import Fuse from "fuse.js";
import React, { useEffect } from "react";
import NavBar from "./molecules/NavBar";
import { TileEntity } from "./types";
import tileJson from "./tiles.json";
import TileEntityList from "./organisms/TileEntityList";

const fuseOptions = {
  // debugging
  // includeScore: true,
  // the keys of the objects to search
  keys: [
    { name: "name", weight: 0.7 },
    { name: "tags", weight: 0.3 },
  ],
  threshold: 0.4,
};

// Assert the structure of the tileJson
const tileData: TileEntity[] = tileJson;

const fuse = new Fuse(tileData, fuseOptions);

export const SearchContext = React.createContext<
  [string | undefined, React.Dispatch<React.SetStateAction<string>> | undefined]
>([undefined, undefined]);

function App(): JSX.Element {
  const [searchVal, setSearchVal] = React.useState("");
  const [searchRes, setSearchRes] = React.useState<TileEntity[]>(tileData);

  useEffect(() => {
    if (searchVal.length > 0) {
      setSearchRes(fuse.search(searchVal).map((res) => res.item));
      return;
    }
    setSearchRes(tileData);
  }, [searchVal]);

  return (
    <div className="App">
      <SearchContext.Provider value={[searchVal, setSearchVal]}>
        <NavBar />
        <TileEntityList list={searchRes} />
      </SearchContext.Provider>
    </div>
  );
}

export default App;
