import { AppProps } from 'next/app';
import Head from 'next/head';
import '@/styles/global.css';
import { DefaultSeo } from 'next-seo';
import nextSeoOptions from '@/api/seoOptions';
import Script from 'next/script';

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      {process.env.NEXT_PUBLIC_GA_TRACKING_ID ? (
        <>
          <Script
            id="google-analytics"
            async
            src={`https://www.googletagmanager.com/gtag/js?id=${process.env.NEXT_PUBLIC_GA_TRACKING_ID}`}
            strategy="lazyOnload"
          />
          <Script id="google-analytics-script" strategy="lazyOnload">
            {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());
          gtag('config', '${process.env.NEXT_PUBLIC_GA_TRACKING_ID}');
          `}
          </Script>
        </>
      ) : null}
      <DefaultSeo {...nextSeoOptions} />
      <Head>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="author" content="RuneMarkers" />
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
        <meta name="theme-color" content="#262624" />
      </Head>
      <Component {...pageProps} />
    </>
  );
}
