import YoutubeEmbed from '../atoms/YouTubeEmbed';

export default function RecommendedGuide({ videoId }: { videoId: string }) {
  return (
    <div
      style={{
        width: `90%`,
      }}
    >
      <h2>Recommended Guide:</h2>
      <YoutubeEmbed videoId={videoId} />
    </div>
  );
}
