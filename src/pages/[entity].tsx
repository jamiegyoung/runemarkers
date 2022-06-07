import { TileEntity } from '@/types';
import { NextSeo } from 'next-seo';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/pages/[entity].module.css';
import StripContainer from '@/components/atoms/StripContainer';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import RecommendedGuide from '@/components/molecules/RecommendedGuide';
import CodeBlock from '@/components/atoms/CodeBlock';
import { tileData } from '@/api/tileJson';

export async function getStaticPaths() {
  return {
    paths: tileData.map((e) => {
      // Super hacky way to get this working, encode all apart from spaces
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
    props: tileData.find((e) => {
      console.log(
        `params.entity`,
        params.entity.toLowerCase(),
        `| e.name`,
        e.name.toLowerCase(),
      );

      return e.safeURI.toLowerCase() === params.entity.toLowerCase();
    }),
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
