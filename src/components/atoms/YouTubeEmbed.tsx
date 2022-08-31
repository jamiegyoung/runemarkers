import styles from '@/components/atoms/YouTubeEmbed.module.css';
import React from 'react';
import LiteYoutubeEmbed from 'react-lite-youtube-embed';
import 'react-lite-youtube-embed/dist/LiteYouTubeEmbed.css';

type YouTubeEmbedProps = {
  videoId: string;
  title: string;
};

export default function YoutubeEmbed({ videoId, title }: YouTubeEmbedProps) {
  return (
    <div className={styles.container}>
      <LiteYoutubeEmbed
        id={videoId}
        title={title}
        poster="hqdefault"
        wrapperClass={[`yt-lite`, styles.customYtLite].join(` `)}
      />
    </div>
  );
}
