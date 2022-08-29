import '@testing-library/jest-dom';
import { event, EventNames } from '@/api/gtag';

const callback = jest.fn();

describe(`gtag`, () => {
  it(`should call gtag with the correct params`, () => {
    window.gtag = callback;
    const eventParam = {
      action: EventNames.copyTileMarkers,
      params: {
        event_category: `test_category`,
        event_label: `test_label`,
        value: 1,
        test_label: `this is a test`,
      },
    };
    event(eventParam);
    expect(callback).toHaveBeenCalledWith(`event`, EventNames.copyTileMarkers, {
      event_category: `test_category`,
      event_label: `test_label`,
      value: 1,
      test_label: `this is a test`,
    });
  });
});
