import YoutubeEmbed from '../atoms/YouTubeEmbed';

export default function RecommendedGuide({
  videoId,
  title,
}: {
  videoId: string;
  title: string;
}) {
  return (
    <div
      style={{
        width: `90%`,
      }}
    >
      <h2>Recommended Guide:</h2>
      <YoutubeEmbed videoId={videoId} title={title} />
    </div>
  );
}
