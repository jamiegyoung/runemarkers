import YoutubeEmbed from '../atoms/YouTubeEmbed';

export default function RecommendedGuide({
  videoId,
  title,
}: {
  videoId: string;
  title: string;
}) {
  return <YoutubeEmbed videoId={videoId} title={title} />;
}
