import { TileEntity } from '@/types';
import tileJson from '@/tiles.json';
import { NextSeo } from 'next-seo';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/pages/[entity].module.css';
import StripContainer from '@/components/atoms/StripContainer';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import RecommendedGuide from '@/components/molecules/RecommendedGuide';
import CodeBlock from '@/components/atoms/CodeBlock';

const tileData: TileEntity[] = tileJson;

export async function getStaticPaths() {
  return {
    paths: tileData.map((e) => {
      return { params: { entity: e.name } };
    }),
    fallback: false,
  };
}

export async function getStaticProps({
  params,
}: {
  params: { entity: string };
}) {
  return {
    props: tileData.find(
      (e) => e.name.toLowerCase() === params.entity.toLowerCase(),
    ),
  };
}

export default function Entity(entity: TileEntity) {
  return (
    <div>
      <NextSeo title="Vorkath Ground Markers" />
      <NavBar />
      <div className={styles.container}>
        <StripContainer>
          <TileEntityCard entity={entity} hideInfoButton />
          {entity.recommendedGuideVideoId ? (
            <RecommendedGuide videoId={entity.recommendedGuideVideoId} />
          ) : null}
          <CodeBlock text={JSON.stringify(entity.tiles)} />
          {entity.source ? (
            <a className={styles.link} href={entity.source}>
              Tiles Source
            </a>
          ) : null}
        </StripContainer>
      </div>
    </div>
  );
}
