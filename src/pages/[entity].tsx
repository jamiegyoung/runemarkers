import { TileEntity } from '@/types';
import { NextSeo } from 'next-seo';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/pages/[entity].module.css';
import StripContainer from '@/components/atoms/StripContainer';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import RecommendedGuide from '@/components/molecules/RecommendedGuide';
import CodeBlock from '@/components/atoms/CodeBlock';
import { getTileData } from '@/api/tiles';

export async function getStaticPaths() {
  return {
    paths: getTileData().map((e) => {
      return {
        params: { entity: e.safeURI },
      };
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
    props: getTileData().find(
      (e) => e.safeURI.toLowerCase() === params.entity.toLowerCase(),
    ),
  };
}

export default function Entity(entity: TileEntity) {
  return (
    <div>
      <NextSeo
        title={`${entity.name} Ground Markers`}
        description={`${entity.name} (${entity.altName}) ground markers for RuneLite. Find and import ground markers for different Oldschool RuneScape activities.`}
      />
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
