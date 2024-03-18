package release

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilterPrefixes(t *testing.T) {

	timestamp := time.Unix(10000000, 0)
	formatted := timestamp.Format(time.RFC3339)

	cases := []struct {
		prefix   string
		expected []Parsed
	}{
		{
			prefix: "foobar-",
			expected: []Parsed{
				{Version: "v1.2.3", Date: timestamp},
			},
		},
		{
			prefix: "some-cool-sdk-",
			expected: []Parsed{
				{Version: "v1.2.4", Date: timestamp},
				{Version: "v1.2.3", Date: timestamp},
			},
		},
		{
			prefix: "",
			expected: []Parsed{
				{Version: "v1.2.3", Date: timestamp},
				{Version: "v1.0.1", Date: timestamp},
				{Version: "v1.0.0-beta.1", Date: timestamp},
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
			assert.Equal(t, c.expected, got)
		})
	}
}
