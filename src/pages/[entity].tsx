import { TileEntity } from '@/types';
import { NextSeo } from 'next-seo';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/pages/[entity].module.css';
import ListContainer from '@/components/atoms/ListContainer';
import TileEntityCard from '@/components/molecules/TileEntityCard';
import YoutubeEmbed from '@/components/atoms/YouTubeEmbed';
import CodeBlock from '@/components/atoms/CodeBlock';
import ContributionFooter from '@/components/atoms/ContributionFooter';
import { getTileData } from '@/api/tiles';
import { defaultImages } from '@/api/seoOptions';
import ListContainerSection from '@/components/molecules/ListContainerSection';
import TilesSource from '@/components/molecules/TilesSource';

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
        title={`${entity.name}${
          entity.subcategory ? ` (${entity.subcategory}) ` : ` `
        }Tile Markers`}
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
        <ListContainer>
          <TileEntityCard entity={entity} hideInfoButton />
          <div className={styles.linkContainer}>
            <a
              href={`https://runelite.net/tile/show/#${Buffer.from(
                JSON.stringify(entity.tiles),
              )
                .toString(`base64`)
                .replaceAll(`=`, ``)}`}
              className={styles.link}
              target="_blank"
              rel="noopener noreferrer"
            >
              View Map on RuneLite
            </a>
            <a
              href={entity.wiki}
              className={styles.link}
              target="_blank"
              rel="noopener noreferrer"
            >
              Wiki Page
            </a>
          </div>
          {entity.recommendedGuideVideoId ? (
            <ListContainerSection title="Recommended Guide">
              <YoutubeEmbed
                videoId={entity.recommendedGuideVideoId}
                title={`${entity.name} guide`}
              />
            </ListContainerSection>
          ) : null}
          <ListContainerSection title="Tile Data">
            <CodeBlock tiles={entity.tiles} />
          </ListContainerSection>

          {entity.source ? (
            <ListContainerSection title="Source">
              <TilesSource source={entity.source} />
            </ListContainerSection>
          ) : null}
        </ListContainer>
      </div>
      <ContributionFooter />
    </>
  );
}
