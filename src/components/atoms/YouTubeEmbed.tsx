import styles from '@/components/atoms/YouTubeEmbed.module.css';

type YouTubeEmbedProps = {
  videoId: string;
};

export default function YoutubeEmbed({ videoId }: YouTubeEmbedProps) {
  return (
    <div className={styles.container}>
      <iframe
        width="512px"
        height="288px"
        src={`https://www.youtube.com/embed/${videoId}`}
        title="YouTube video player"
        frameBorder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowFullScreen
      />
    </div>
  );
}
