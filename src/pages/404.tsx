import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/404.module.css';
import { NextSeo } from 'next-seo';
import Head from 'next/head';
import Image from 'next/image';
import Link from 'next/link';

export default function FourOhFour() {
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
            src={`https://thumbs.gfycat.com/SecondaryUncommonKestrel.webp`}
            alt="404 gif"
            width={270}
            height={490}
            layout="intrinsic"
            objectFit="contain"
            priority={true}
          />
        </div>
        <Link href="/">
          <span className={styles.home}>Go Home</span>
        </Link>
      </div>
    </>
  );
}
