import { AppProps } from 'next/app';
import Head from 'next/head';
import '@/styles/global.css';
import { DefaultSeo } from 'next-seo';
import nextSeoOptions from '@/api/seoOptions';

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <DefaultSeo {...nextSeoOptions} />
      <Head>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="author" content="RuneMarkers" />
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="theme-color" content="#262624" />
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href="/apple-touch-icon.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href="/favicon-32x32.png"
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href="/favicon-16x16.png"
        />
        <link rel="manifest" href="/site.webmanifest" />
      </Head>
      <Component {...pageProps} />
    </>
  );
}
