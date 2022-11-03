import Search from './Search';
import styles from '@/components/molecules/ListOperations.module.css';
import { ViewFormat } from '../organisms/TileEntities';
import LayoutButton from '../atoms/LayoutButton';

export default function ListOperations({
  viewFormat,
  setViewFormat,
}: {
  viewFormat: ViewFormat;
  setViewFormat: (format: ViewFormat) => void;
}) {
  const getNextListValue = () =>
    (viewFormat + 1) % (Object.keys(ViewFormat).length / 2);

  return (
    <div
      className={[
        styles.listOperationsContainer,
        styles[ViewFormat[viewFormat].toLowerCase()],
      ].join(` `)}
    >
      <Search />
      <LayoutButton
        viewFormat={viewFormat}
        onClick={() => setViewFormat(getNextListValue())}
      />
    </div>
  );
}
