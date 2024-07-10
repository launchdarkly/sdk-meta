package products

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestProductSanityChecks(t *testing.T) {
	t.Run("names", func(t *testing.T) {
		assert.Equal(t, "Node.js Server SDK", Names["node-server"])
	})

	t.Run("repos", func(t *testing.T) {
		assert.Equal(t, "launchdarkly/js-core", Repos["node-server"].Github)
	})

	t.Run("languages", func(t *testing.T) {
		assert.Equal(t, []string{"JavaScript", "TypeScript"}, Languages["node-server"])
	})

	t.Run("types", func(t *testing.T) {
		assert.Equal(t, ServerSideType, Types["node-server"])
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

		assert.Equal(t, firstNodeReleaseDate, firstRelease.Date)
		require.NotNil(t, firstRelease.EOL)
		assert.Equal(t, firstNodeReleaseEOL, *firstRelease.EOL)

		latestRelease := Releases["node-server"].Latest()
		assert.GreaterOrEqual(t, latestRelease.Major, 9)
		assert.GreaterOrEqual(t, latestRelease.Minor, 4)
		assert.Nil(t, latestRelease.EOL)
	})

}
