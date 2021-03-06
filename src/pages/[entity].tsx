import { TileEntity } from '@/types';
import { NextSeo } from 'next-seo';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/pages/[entity].module.css';
import StripContainer from '@/components/atoms/StripContainer';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import RecommendedGuide from '@/components/molecules/RecommendedGuide';
import CodeBlock from '@/components/atoms/CodeBlock';
import ContributionFooter from '@/components/atoms/ContributionFooter';
import { getTileData } from '@/api/tiles';
import { defaultImages } from '@/api/seoOptions';

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
    <>
      <NextSeo
        title={`${entity.name} Tile Markers`}
        description={`${entity.name}${
          entity.altName ? ` / ${entity.altName}` : ``
        } tile markers for RuneLite. Find and import tile markers for different Oldschool RuneScape activities.`}
        openGraph={{
          images: [
            {
              url: entity.thumbnail,
              alt: `${entity.name} tile markers`,
            },
            ...defaultImages,
          ],
        }}
      />
      <NavBar />
      <div className={styles.container}>
        <StripContainer>
          <TileEntityCard entity={entity} hideInfoButton />
          <div className={styles.linkContainer}>
            <a
              href={`https://runelite.net/tile/show/#${Buffer.from(
                JSON.stringify(entity.tiles),
              )
                .toString(`base64`)
                .replaceAll(`=`, ``)}`}
              className={styles.link}
            >
              View Tiles on RuneLite
            </a>
            <a href={entity.wiki} className={styles.link}>
              Wiki Page
            </a>
          </div>
          {entity.recommendedGuideVideoId ? (
            <RecommendedGuide
              videoId={entity.recommendedGuideVideoId}
              title={`${entity.name} guide`}
            />
          ) : null}
          <CodeBlock text={JSON.stringify(entity.tiles)} />
          {entity.source ? (
            <a className={styles.link} href={entity.source}>
              Tiles Source
            </a>
          ) : null}
        </StripContainer>
      </div>
      <ContributionFooter />
    </>
  );
}
