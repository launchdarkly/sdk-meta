// Jest test that drives the snippet's App.tsx + src/welcome.tsx end to
// end. Renders the React tree, waits up to 30s for the SDK to deliver
// the flag, then asserts that the rendered text contains the canonical
// EXAM-HELLO line.
//
// The SDK's default streaming transport uses XMLHttpRequest under the
// hood (RNEventSource), which doesn't exist in Node. Force polling
// mode (which uses fetch — built into Node 18+) so the test can talk
// to LaunchDarkly's polling endpoint and resolve the flag.
import React from 'react';
import {render, waitFor, screen} from '@testing-library/react-native';

jest.mock('@launchdarkly/react-native-client-sdk', () => {
  const actual = jest.requireActual('@launchdarkly/react-native-client-sdk');
  class PollingLDClient extends actual.ReactNativeLDClient {
    constructor(key: string, autoEnv: any, options: any = {}) {
      super(key, autoEnv, {...options, initialConnectionMode: 'polling'});
    }
  }
  return {...actual, ReactNativeLDClient: PollingLDClient};
});

// eslint-disable-next-line import/first
import App from '../App';

jest.setTimeout(60_000);

// Walks the rendered tree and concatenates every Text child into a
// single flat string. Needed because react-native's render output is a
// nested object — a regex looking for "feature flag evaluates to true"
// only matches when the words are adjacent.
function flattenText(node: any): string {
  if (node == null) return '';
  if (typeof node === 'string') return node;
  if (Array.isArray(node)) return node.map(flattenText).join('');
  if (typeof node === 'object' && node.children) return flattenText(node.children);
  return '';
}

test('flag evaluates to true', async () => {
  render(<App />);
  await waitFor(
    () => {
      const text = flattenText(screen.toJSON());
      expect(text).toMatch(/feature flag evaluates to true/i);
    },
    {timeout: 30_000, interval: 500},
  );
  // Print the flat text so the validator harness's grep on
  // "feature flag evaluates to [Tt]rue" matches.
  console.log(flattenText(screen.toJSON()));
});
