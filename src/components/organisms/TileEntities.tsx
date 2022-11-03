import {
  HeightScrollTop,
  TileEntity,
  TileEntityCardCompactProps,
  TileEntityCardProps,
} from '@/types';
import styles from '@/components/organisms/TileEntities.module.css';
import ListContainer from '@/components/atoms/ListContainer';
import {
  ReactNode,
  RefObject,
  useCallback,
  useContext,
  useEffect,
  useRef,
  useState,
} from 'react';
import { SearchContext } from '@/pages';
import GridContainer from '../atoms/GridContainer';
import ListOperations from '../molecules/ListOperations';
import useWindowDimensions from '@/hooks/useWindowDimensions';
import TileEntityCard from '../molecules/TileEntityCard';
import TileEntityCardCompact from '../molecules/TileEntityCardCompact';

type TileEntitiesProps = {
  list: TileEntity[];
};

export enum ViewFormat {
  List,
  Grid,
}

export default function TileEntities({ list }: TileEntitiesProps) {
  const ref = useRef<HTMLDivElement>(null);
  const [searchVal] = useContext(SearchContext);
  const [viewFormat, _setViewFormat] = useState<ViewFormat | undefined>(
    undefined,
  );
  const [previousScrollInfo, setPreviousScrollInfo] = useState<HeightScrollTop>(
    { height: 0, scrollTop: 0 },
  );
  const { width } = useWindowDimensions();

  const setViewFormat = useCallback(
    (format: ViewFormat) => {
      if (format == undefined || format === viewFormat) return;
      localStorage.setItem(`viewFormat`, format.toString());
      if (ref.current) {
        setPreviousScrollInfo({
          height: ref.current.scrollHeight,
          scrollTop: ref.current.scrollTop,
        });
      }
      _setViewFormat(format);
      localStorage.setItem(`viewFormat`, format.toString());
    },
    [viewFormat],
  );

  useEffect(() => {
    if (!ref.current || !ref.current.scrollTo || !previousScrollInfo) return;
    ref.current.scrollTo(
      0,
      previousScrollInfo.scrollTop /
        (previousScrollInfo.height / ref.current.scrollHeight),
    );
  }, [viewFormat, previousScrollInfo]);

  useEffect(() => {
    if (width < 600) {
      setViewFormat(ViewFormat.List);
    }
  }, [setViewFormat, width]);

  useEffect(() => {
    const viewFormat = localStorage.getItem(`viewFormat`);
    if (viewFormat !== undefined) {
      setViewFormat(Number(viewFormat));
      return;
    }
    setViewFormat(ViewFormat.List);
  }, [setViewFormat]);

  useEffect(() => {
    if (ref.current) {
      ref.current.scrollTop = 0;
    }
  }, [searchVal, ref]);

  function entityContainer(
    Container: ({
      children,
      customRef,
    }: {
      children: ReactNode;
      customRef?: RefObject<HTMLDivElement> | undefined;
    }) => JSX.Element,
    Card: ({
      entity,
    }: TileEntityCardProps | TileEntityCardCompactProps) => JSX.Element,
  ) {
    return (
      <Container customRef={ref}>
        {list.length > 0 ? (
          list.map((i) => <Card key={i.safeURI} entity={i} />)
        ) : (
          <p style={{ textAlign: `center` }}>No results found</p>
        )}
      </Container>
    );
  }

  return (
    <div className={styles.container}>
      <p
        style={{
          padding: `0`,
          color: `#C7C7C7`,
          textAlign: `center`,
          minWidth: 280,
          marginBottom: 15,
          maxWidth: 600,
          width: `80%`,
        }}
      >
        Make sure the &quot;Ground Markers&quot; plugin is enabled on RuneLite
        with &quot;Remember color per tile&quot; enabled. Copy the tile markers
        to your clipboard, right-click the world map button in-game and select
        &quot;import&quot;.
      </p>
      {viewFormat !== undefined ? (
        <ListOperations viewFormat={viewFormat} setViewFormat={setViewFormat} />
      ) : null}
      {viewFormat === ViewFormat.List
        ? entityContainer(ListContainer, TileEntityCard)
        : null}
      {viewFormat === ViewFormat.Grid
        ? entityContainer(GridContainer, TileEntityCardCompact)
        : null}
      {viewFormat === undefined ? (
        <p style={{ textAlign: `center` }}>Loading...</p>
      ) : null}
    </div>
  );
}
