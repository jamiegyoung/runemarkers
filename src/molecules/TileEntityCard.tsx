import Button from "../atoms/Button";
import InfoButton from "../atoms/InfoButton";
import { TileEntity } from "../types";
import styles from "./TileEntityCard.module.css";
import { LazyLoadImage } from "react-lazy-load-image-component";
import useWindowDimensions from '../hooks/useWindowDimensions';

type TileEntityCardProps = {
  entity: TileEntity;
};

export default function TileEntityCard({ entity }: TileEntityCardProps) {
  const { width } = useWindowDimensions();
  return (
    <div className={styles.card}>
      <LazyLoadImage
        className={styles.image}
        width={width > 450 ? 200 : 100}
        src={entity.thumbnail}
        alt={`${entity.name}.png`}
      />
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <h1>{entity.name}</h1>
          <p>{entity.tiles.length} tile markers</p>
        </div>
        <div className={styles.tileInteraction}>
          <Button
            style={{ marginRight: 10, padding: "10px 40px" }}
            onClick={() => {
              navigator.clipboard.writeText(JSON.stringify(entity.tiles));
            }}
          >
            Copy
          </Button>
          <InfoButton />
        </div>
      </div>
    </div>
  );
}
