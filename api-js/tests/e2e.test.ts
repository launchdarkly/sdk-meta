import { Names, Repos, Types, Type, Popularity, Languages, Releases, ReleaseHelpers } from '../src/SDKMeta';
import { UserAgents, UserAgentHelpers } from '../src/SDKMeta';

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
    expect(Types['node-server']).toBe(Type.ServerSide);
    expect(Types['node-server']).toBe('server-side');
});

test('popularity', () => {
    expect(Popularity['node-server']).toBe(2);
});

test('releases', () => {
    const firstNodeReleaseDate = new Date("2015-05-13T16:55:00Z");
    const firstNodeReleaseEOL = new Date("2016-09-12T00:00:00Z");

    expect(Releases['node-server'].length).toBeGreaterThanOrEqual(1);

    const firstRelease = ReleaseHelpers.Earliest(Releases['node-server']);
    expect(firstRelease.Major).toBe(1);
    expect(firstRelease.Minor).toBe(0);
    expect(ReleaseHelpers.IsLatest(firstRelease)).toBe(false);

    expect(firstRelease.Date).toEqual(firstNodeReleaseDate);
    expect(firstRelease.EOL).not.toBeNull();
    expect(firstRelease.EOL).toEqual(firstNodeReleaseEOL);

    const latestRelease = ReleaseHelpers.Latest(Releases['node-server']);
    expect(latestRelease.Major).toBeGreaterThanOrEqual(9);
    expect(latestRelease.Minor).toBeGreaterThanOrEqual(4);
    expect(latestRelease.EOL).toBeNull();
    expect(ReleaseHelpers.IsLatest(latestRelease)).toBe(true);
})

test('eol calculations', () => {
    const releases = Releases['node-server'];
    const earliest = ReleaseHelpers.Earliest(releases);
    const latest = ReleaseHelpers.Latest(releases);
    const earliestEOL = new Date("2016-09-12T00:00:00Z");

    // Checking that the latest release is not yet EOL
    expect(ReleaseHelpers.IsEOL(latest, new Date())).toBe(false);
    // Checking that the earliest release becomes EOL if we pass in a "current" date of its EOL + 1 second
    expect(ReleaseHelpers.IsEOL(earliest, new Date(earliestEOL.getTime() + 1000))).toBe(true);

    // Check the "approaching EOL" logic for the earliest release by passing in different values of "current" date.
    const minute = 60 * 1000;
    const hour = 60 * minute;
    const hour_and_1_minute = 61 * minute;
    const fifty_nine_minutes = 59 * minute;
    const thirty_minutes = 30 * minute;

    expect(ReleaseHelpers.IsApproachingEOL(earliest, new Date(earliestEOL.getTime() - hour_and_1_minute), hour)).toBe(false);
    expect(ReleaseHelpers.IsApproachingEOL(earliest, new Date(earliestEOL.getTime() - hour), hour)).toBe(false);
    expect(ReleaseHelpers.IsApproachingEOL(earliest, new Date(earliestEOL.getTime() - fifty_nine_minutes), hour)).toBe(true);
    expect(ReleaseHelpers.IsApproachingEOL(earliest, new Date(earliestEOL.getTime() - thirty_minutes), hour)).toBe(true);
    expect(ReleaseHelpers.IsApproachingEOL(earliest, new Date(earliestEOL.getTime() - minute), hour)).toBe(true);
})

test('user agents', () => {
    // Test basic user agent data structure
    expect(UserAgents['node-server']).toBeDefined();
    expect(UserAgents['node-server'].userAgents).toContain('NodeJSClient');

    // Test SDK with both user agents and wrapper names
    expect(UserAgents['react-native']).toBeDefined();
    expect(UserAgents['react-native'].userAgents).toContain('ReactNativeClient');
    expect(UserAgents['react-native'].wrapperNames).toContain('react-native-client');
});

describe('UserAgentHelpers.getSDKNameByUserAgentOrWrapper', () => {
    test('finds SDK by wrapper name', () => {
        const name = UserAgentHelpers.getSDKNameByUserAgentOrWrapper('react-native-client');
        expect(name).toBe('React Native SDK');
    });

    test('finds SDK by user agent', () => {
        const name = UserAgentHelpers.getSDKNameByUserAgentOrWrapper('NodeJSClient');
        expect(name).toBe('Node.js Server SDK');
    });

    test('returns undefined for unknown identifier', () => {
        const name = UserAgentHelpers.getSDKNameByUserAgentOrWrapper('UnknownIdentifier');
        expect(name).toBeUndefined();
    });

    test('prioritizes wrapper names over user agents', () => {
        // In case there's ever a wrapper name that matches a user agent from another SDK,
        // we should ensure wrapper names are checked first
        const name = UserAgentHelpers.getSDKNameByUserAgentOrWrapper('react-native-client');
        expect(name).toBe('React Native SDK');
    });

    test('finds edge SDK by user agent', () => {
        const name = UserAgentHelpers.getSDKNameByUserAgentOrWrapper('CloudflareEdgeSDK');
        expect(name).toBe('Cloudflare SDK');
    });
});
