import { Names, Repos, Types, Popularity, Languages } from '../src/SDKMeta';

test('names', () => {
    expect(Names['node-server']).toBe('Node.js Server SDK');
});

test('repos', () => {
    expect(Repos['node-server'].github).toBe('launchdarkly/js-core');
});

test('languages', () => {
    expect(Languages['node-server']).toEqual(['JavaScript', 'TypeScript']);
});

test('types', () => {
    expect(Types['node-server']).toBe('server-side');
});

test('popularity', () => {
    expect(Popularity['node-server']).toBe(2);
});
