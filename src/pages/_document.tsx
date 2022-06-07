import { Html, Head, Main, NextScript } from 'next/document';

export default function Document() {
  return (
    <Html>
      <Head>
        <meta name="author" content="RuneMarkers" />
        <meta property="og:site_name" content="RuneMarkers" />
        <meta property="og:type" content="website" />
        <meta property="og:url" content="https://runemarkers.net" />
        <meta property="og:image" content="/logo1024.png" />
        <meta charSet="utf-8" />
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
        <link rel="manifest" href="/site.webmanifest.json" />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
