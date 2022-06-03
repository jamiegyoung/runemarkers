import TileEntityCard from "../molecules/TileEntityCard";
import { TileEntity } from "../types";
import styles from "./TileEntityList.module.css";

type TileEntityListProps = {
  list: TileEntity[];
};

export default function TileEntityList({ list }: TileEntityListProps) {
  console.log(list);
  return (
    <div className={styles.container}>
      <p
        style={{
          padding: "0 20px",
          color: "#C7C7C7",
          maxWidth: "500px",
          textAlign: "center",
        }}
      >
        Make sure the “Gound Markers” plugin is enabled in runelite, copy the
        tile markers to your clipboard, right click the world map on runelite
        and then select “import”
      </p>
      <div className={styles.innerContainer}>
        {list.map((i) => (
          <TileEntityCard key={i.name} entity={i} />
        ))}
      </div>
    </div>
  );
}
