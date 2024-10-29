const sdkMeta = require('../dist/SDKMeta.cjs');

test('names', () => {
    expect(sdkMeta.Names['node-server']).toBe('Node.js Server SDK');
});

test('repos', () => {
    expect(sdkMeta.Repos['node-server'].github).toBe('launchdarkly/js-core');
});

test('languages', () => {
    expect(sdkMeta.Languages['node-server']).toEqual(['JavaScript', 'TypeScript']);
});

test('types', () => {
    expect(sdkMeta.Types['node-server']).toBe('server-side');
});

test('popularity', () => {
    expect(sdkMeta.Popularity['node-server']).toBe(2);
});
