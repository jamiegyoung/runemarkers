import Button from "../atoms/Button";
import InfoButton from "../atoms/InfoButton";
import { TileEntity } from "../types";
import styles from "./TileEntityCard.module.css";
import { LazyLoadImage } from "react-lazy-load-image-component";
import useWindowDimensions from "../hooks/useWindowDimensions";
import copy from "copy-to-clipboard";

type TileEntityCardProps = {
  entity: TileEntity;
};

export default function TileEntityCard({ entity }: TileEntityCardProps) {
  const { width } = useWindowDimensions();
  return (
    <div className={styles.card}>
      <a href={entity.wiki} target="_blank" rel="noopener noreferrer"
        className={styles.pictureLink}
      >
        <LazyLoadImage
          className={styles.image}
          width={width > 460 ? 200 : 100}
          src={entity.thumbnail}
          alt={`${entity.name}.png`}
        />
      </a>
      <div className={styles.cardData}>
        <div className={styles.entityInfo}>
          <h1>{entity.name}</h1>
          <p>{entity.tiles.length} tile markers</p>
        </div>
        <div className={styles.tileInteraction}>
          <Button
            style={{ marginRight: width < 330 ? 0 : 10, padding: width < 400 ? "10px 20px" : "10px 40px" }}
            onClick={() => {
              copy(JSON.stringify(entity.tiles));
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
