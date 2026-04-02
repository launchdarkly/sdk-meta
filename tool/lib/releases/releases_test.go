package releases

import (
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestFilterPrefixes(t *testing.T) {

	timestamp := time.Unix(10000000, 0).UTC()
	formatted := timestamp.Format(time.RFC3339)

	cases := []struct {
		prefix   string
		expected []Parsed
	}{
		{
			prefix: "foobar-",
			expected: []Parsed{
				{Version: semver.MustParse("v1.2.3"), Date: timestamp},
			},
		},
		{
			prefix: "some-cool-sdk-",
			expected: []Parsed{
				{Version: semver.MustParse("v1.2.3"), Date: timestamp},

				{Version: semver.MustParse("v1.2.4"), Date: timestamp},
			},
		},
		{
			prefix: "",
			expected: []Parsed{
				{Version: semver.MustParse("v1.2.3"), Date: timestamp},
				{Version: semver.MustParse("v1.0.1"), Date: timestamp},
				{Version: semver.MustParse("v1.0.0-beta.1"), Date: timestamp},
			},
		},
	}

	releases := []Raw{
		{Tag: "foobar-v1.2.3", Date: formatted},
		{Tag: "some-cool-sdk-v1.2.3", Date: formatted},
		{Tag: "some-cool-sdk-v1.2.4", Date: formatted},
		{Tag: "v1.2.3", Date: formatted},
		{Tag: "1.0.1", Date: formatted},
		{Tag: "1.0.0-beta.1", Date: formatted},
	}

	for _, c := range cases {
		t.Run(c.prefix, func(t *testing.T) {
			got, err := Filter(releases, c.prefix)
			assert.Nil(t, err)
			for i := 0; i < len(got); i++ {
				assert.Equal(t, c.expected[i].Version.String(), got[i].Version.String())
				assert.Equal(t, c.expected[i].Date, got[i].Date)
			}
		})
	}
}

func TestFilterMulti(t *testing.T) {
	timestamp := time.Unix(10000000, 0).UTC()
	formatted := timestamp.Format(time.RFC3339)

	releases := []Raw{
		{Tag: "my-sdk-v2.0.0", Date: formatted},
		{Tag: "my-sdk-v1.5.0", Date: formatted},
		{Tag: "1.4.0", Date: formatted},
		{Tag: "1.3.0", Date: formatted},
		{Tag: "1.0.0-beta.1", Date: formatted},
	}

	t.Run("matches both bare and prefixed tags", func(t *testing.T) {
		got, err := FilterMulti(releases, []string{"", "my-sdk-"})
		assert.Nil(t, err)
		assert.Len(t, got, 5)
		assert.Equal(t, "2.0.0", got[0].Version.String())
		assert.Equal(t, "1.5.0", got[1].Version.String())
		assert.Equal(t, "1.4.0", got[2].Version.String())
		assert.Equal(t, "1.3.0", got[3].Version.String())
		assert.Equal(t, "1.0.0-beta.1", got[4].Version.String())
	})

	t.Run("deduplicates overlapping versions", func(t *testing.T) {
		duped := make([]Raw, len(releases), len(releases)+1)
		copy(duped, releases)
		duped = append(duped, Raw{Tag: "v1.4.0", Date: formatted})
		got, err := FilterMulti(duped, []string{"", "my-sdk-"})
		assert.Nil(t, err)
		assert.Len(t, got, 5)
	})

	t.Run("single prefix behaves like Filter", func(t *testing.T) {
		got, err := FilterMulti(releases, []string{"my-sdk-"})
		assert.Nil(t, err)
		assert.Len(t, got, 2)
		assert.Equal(t, "2.0.0", got[0].Version.String())
		assert.Equal(t, "1.5.0", got[1].Version.String())
	})
}
