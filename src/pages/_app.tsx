import { AppProps } from 'next/app';
import Head from 'next/head';
import '@/styles/global.css';

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <Component {...pageProps} />;
    </>
  );
}
