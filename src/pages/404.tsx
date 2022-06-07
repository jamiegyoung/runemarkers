import Button from '@/components/atoms/Button';
import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/404.module.css';
import { NextSeo } from 'next-seo';
import Head from 'next/head';
import Image from 'next/image';
import { useState } from 'react';

export default function FourOhFour() {
  const [isAnimating, setIsAnimating] = useState(false);
  return (
    <>
      <Head>
        <link href="runescape_uf.ttf" as="font" type="font/woff2" />
        <NextSeo title="404" description="Page not found" />
      </Head>
      <NavBar />
      <div className={styles.container}>
        <h1 className={styles.header}>404</h1>
        <div className={styles.imageContainer}>
          <Image
            className={styles.image}
            src={
              isAnimating
                ? `https://thumbs.gfycat.com/DefinitiveLimpArawana.webp`
                : `https://thumbs.gfycat.com/SecondaryUncommonKestrel.webp`
            }
            alt="404 gif"
            width={270}
            height={490}
            layout="intrinsic"
            objectFit="contain"
            sizes="50vw"
          />
        </div>
        <Button
          className={styles.home}
          onClick={() => {
            setIsAnimating(true);
            setTimeout(() => {
              window.location.replace(`/`);
            }, 1600);
          }}
        >
          Go Home
        </Button>
      </div>
    </>
  );
}
