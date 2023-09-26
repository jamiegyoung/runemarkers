jest.mock('@vercel/analytics/react', () => ({
  Analytics: () => null,
}));

