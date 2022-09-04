import NavBar from '@/components/molecules/NavBar';
import styles from '@/styles/404.module.css';
import { NextSeo } from 'next-seo';
import Head from 'next/head';
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
        <Link href="/">
          <span className={styles.home}>Go Home</span>
        </Link>
      </div>
    </>
  );
}
