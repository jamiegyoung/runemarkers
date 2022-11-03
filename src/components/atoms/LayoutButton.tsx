import { ViewFormat } from '../organisms/TileEntities';
import styles from '@/components/atoms/LayoutButton.module.css';

export default function LayoutButton({
  viewFormat: viewFormat,
  onClick,
}: {
  viewFormat: ViewFormat;
  onClick: () => void;
}) {
  return (
    <button
      className={styles.button}
      onClick={() => onClick()}
      type="button"
      aria-label="Change layout"
      title="Change layout"
    >
      <div className={styles.innerButton}>
        <div
          className={[
            styles.buttonSection,
            // lmao
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
        <div
          className={[
            styles.buttonSection,
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
        <div
          className={[
            styles.buttonSection,
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
        <div
          className={[
            styles.buttonSection,
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
        <div
          className={[
            styles.buttonSection,
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
        <div
          className={[
            styles.buttonSection,
            styles[ViewFormat[viewFormat].toLowerCase()],
          ].join(` `)}
        />
      </div>
    </button>
  );
}
