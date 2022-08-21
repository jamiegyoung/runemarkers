import { AppProps } from 'next/app';
import Head from 'next/head';
import '@/styles/global.css';
import { DefaultSeo } from 'next-seo';
import nextSeoOptions from '@/api/seoOptions';
import Script from 'next/script';
import { createContext, MutableRefObject, useContext, useRef } from 'react';

const GtagContext = createContext<null | MutableRefObject<
  (
    type: string,
    eventName: string,
    options: {
      tiles_copied: string;
    },
  ) => void
>>(null);

export const useGtagContext = () => useContext(GtagContext);

export default function MyApp({ Component, pageProps }: AppProps) {
  const gtagRef = useRef<any>();

  return (
    <>
      {process.env.NEXT_PUBLIC_GA_TRACKING_ID ? (
        <>
          <Script
            id="google-analytics"
            async
            src={`https://www.googletagmanager.com/gtag/js?id=${process.env.NEXT_PUBLIC_GA_TRACKING_ID}`}
            strategy="afterInteractive"
          />
          <Script
            id="google-analytics-script"
            strategy="afterInteractive"
            onReady={() => {
              // eslint-disable-next-line @typescript-eslint/ban-ts-comment
              // @ts-ignore
              gtagRef.current = window.gtag;
            }}
          >
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
      <GtagContext.Provider value={gtagRef}>
        <Component {...pageProps} />
      </GtagContext.Provider>
    </>
  );
}
