module.exports = {
  reactStrictMode: true,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'oldschool.runescape.wiki',
        pathname: '/images/**'
      }
    ]
  },
  swcMinify: true,
  output: 'standalone',
};
