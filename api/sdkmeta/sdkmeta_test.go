package sdkmeta

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductSanityChecks(t *testing.T) {
	t.Run("names", func(t *testing.T) {
		assert.Equal(t, "Node.js Server SDK", Names["node-server"])
	})

	t.Run("repos", func(t *testing.T) {
		assert.Equal(t, "launchdarkly/js-core", Repos["node-server"].GitHub)
	})

	t.Run("languages", func(t *testing.T) {
		assert.Equal(t, []string{"JavaScript", "TypeScript"}, Languages["node-server"])
	})

	t.Run("types", func(t *testing.T) {
		assert.Equal(t, ServerSideType, Types["node-server"])
	})

	t.Run("popularity", func(t *testing.T) {
		assert.Equal(t, 2, Popularity["node-server"])
	})

	t.Run("releases", func(t *testing.T) {
		firstNodeReleaseDate, err := time.Parse(time.RFC3339, "2015-05-13T16:55:00Z")
		require.NoError(t, err)
		firstNodeReleaseEOL, err := time.Parse(time.RFC3339, "2016-09-12T00:00:00Z")
		require.NoError(t, err)

		require.GreaterOrEqual(t, len(Releases["node-server"]), 1, "Expected at least one release for node-server")

		firstRelease := Releases["node-server"].Earliest()
		assert.Equal(t, 1, firstRelease.Major)
		assert.Equal(t, 0, firstRelease.Minor)
		assert.False(t, firstRelease.IsLatest())

		assert.Equal(t, firstNodeReleaseDate, firstRelease.Date)
		require.NotNil(t, firstRelease.EOL)
		assert.Equal(t, firstNodeReleaseEOL, *firstRelease.EOL)

		latestRelease := Releases["node-server"].Latest()
		assert.GreaterOrEqual(t, latestRelease.Major, 9)
		assert.GreaterOrEqual(t, latestRelease.Minor, 4)
		assert.Nil(t, latestRelease.EOL)
		assert.True(t, latestRelease.IsLatest())
	})
}

func TestEOLCalculations(t *testing.T) {
	releases := Releases["node-server"]
	earliest := releases.Earliest()
	latest := releases.Latest()
	earliestEOL := time.Date(2016, 9, 12, 0, 0, 0, 0, time.UTC)

	t.Run("is eol", func(t *testing.T) {
		assert.False(t, latest.IsEOL(time.Now()))
		assert.True(t, earliest.IsEOL(earliestEOL.Add(time.Second)))
	})

	t.Run("is approaching eol", func(t *testing.T) {
		assert.False(t, earliest.IsApproachingEOL(earliestEOL.Add(-61*time.Minute), time.Hour))
		assert.False(t, earliest.IsApproachingEOL(earliestEOL.Add(-60*time.Minute), time.Hour))
		assert.True(t, earliest.IsApproachingEOL(earliestEOL.Add(-59*time.Minute), time.Hour))
		assert.True(t, earliest.IsApproachingEOL(earliestEOL.Add(-30*time.Minute), time.Hour))
		assert.True(t, earliest.IsApproachingEOL(earliestEOL.Add(-1*time.Minute), time.Hour))
	})
}

func TestUserAgentsAndWrappers(t *testing.T) {
	t.Run("user agents map contains expected data", func(t *testing.T) {
		nodeInfo := UserAgents["node-server"]
		assert.Contains(t, nodeInfo.UserAgents, "NodeJSClient")

		reactNativeInfo := UserAgents["react-native"]
		assert.Contains(t, reactNativeInfo.UserAgents, "ReactNativeClient")
		assert.Contains(t, reactNativeInfo.WrapperNames, "react-native-client")
	})

	t.Run("GetSDKNameByWrapperOrUserAgent finds by wrapper", func(t *testing.T) {
		name, found := GetSDKNameByWrapperOrUserAgent("react-native-client")
		assert.True(t, found)
		assert.Equal(t, "React Native SDK", name)
	})

	t.Run("GetSDKNameByWrapperOrUserAgent finds by user agent", func(t *testing.T) {
		name, found := GetSDKNameByWrapperOrUserAgent("NodeJSClient")
		assert.True(t, found)
		assert.Equal(t, "Node.js Server SDK", name)
	})

	t.Run("GetSDKNameByWrapperOrUserAgent returns false for unknown identifier", func(t *testing.T) {
		name, found := GetSDKNameByWrapperOrUserAgent("UnknownIdentifier")
		assert.False(t, found)
		assert.Empty(t, name)
	})
}

func TestUserAgentMapResolveSDKID(t *testing.T) {
	t.Run("finds SDK ID by user agent", func(t *testing.T) {
		id, found := UserAgents.ResolveSDKID("NodeJSClient")
		assert.True(t, found)
		assert.Equal(t, "node-server", id)
	})

	t.Run("finds SDK ID by wrapper name", func(t *testing.T) {
		id, found := UserAgents.ResolveSDKID("ElectronClient")
		assert.True(t, found)
		assert.Equal(t, "electron", id)
	})

	t.Run("returns the ID that Names is keyed by", func(t *testing.T) {
		id, found := UserAgents.ResolveSDKID("RokuClient")
		require.True(t, found)
		assert.Equal(t, "roku", id)
		assert.Equal(t, "Roku SDK", Names[id])
	})

	t.Run("returns false for an unknown identifier", func(t *testing.T) {
		id, found := UserAgents.ResolveSDKID("NotARealClient")
		assert.False(t, found)
		assert.Empty(t, id)
	})

	t.Run("agrees with GetSDKNameByWrapperOrUserAgent", func(t *testing.T) {
		for _, identifier := range []string{"NodeJSClient", "ElectronClient", "GoClient"} {
			id, foundID := UserAgents.ResolveSDKID(identifier)
			name, foundName := GetSDKNameByWrapperOrUserAgent(identifier)
			require.Equal(t, foundID, foundName)
			assert.Equal(t, Names[id], name)
		}
	})
}

func TestAISDKIdentifierMapResolveSDKID(t *testing.T) {
	t.Run("every identifier resolves to a known SDK", func(t *testing.T) {
		for sdkID, identifiers := range AISDKIdentifiers {
			assert.Contains(t, Names, sdkID, "AI SDK %s is missing from names", sdkID)
			assert.Equal(t, AIType, Types[sdkID], "AI SDK %s should be typed ai", sdkID)
			for _, identifier := range identifiers {
				resolved, found := AISDKIdentifiers.ResolveSDKID(identifier.Name, identifier.Language)
				assert.True(t, found)
				assert.Equal(t, sdkID, resolved)
			}
		}
	})

	t.Run("returns false for an unknown name", func(t *testing.T) {
		id, found := AISDKIdentifiers.ResolveSDKID("not-a-real-package", "python")
		assert.False(t, found)
		assert.Empty(t, id)
	})

	t.Run("language is required to disambiguate", func(t *testing.T) {
		// The Python and Ruby AI SDKs report the same package name, so a lookup
		// without the correct language must not resolve.
		id, found := AISDKIdentifiers.ResolveSDKID("launchdarkly-server-sdk-ai", "")
		assert.False(t, found)
		assert.Empty(t, id)
	})
}

// Language values in ai_sdk_identifiers.json are produced by lowercasing an SDK's
// declared language, so every value in the language vocabulary must survive the
// schema's pattern. C#, C++, and Objective-C are the ones that catch a naive [a-z]+.
func TestAISDKIdentifierLanguagePattern(t *testing.T) {
	pattern := regexp.MustCompile(`^[a-z0-9+#.-]+$`)
	seen := map[string]bool{}
	for _, languages := range Languages {
		for _, language := range languages {
			seen[language] = true
		}
	}
	require.NotEmpty(t, seen)
	for language := range seen {
		assert.Regexp(t, pattern, strings.ToLower(language),
			"lowercased %q must match the ai_sdk_identifiers language pattern", language)
	}
}
