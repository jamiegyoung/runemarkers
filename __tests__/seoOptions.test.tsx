import '@testing-library/jest-dom';
import { render } from '@testing-library/react';

import nextSeoOptions from '@/api/seoOptions';
import { DefaultSeo } from 'next-seo';

describe(`nextSeoOptions`, () => {
  it(`should render correctly`, () => {
    const { container } = render(<DefaultSeo {...nextSeoOptions} />);
    expect(container).toMatchSnapshot();
  });
});
