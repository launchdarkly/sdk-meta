package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEffectivePrefixes(t *testing.T) {
	tests := []struct {
		name     string
		metadata metadataV1
		expected []string
	}{
		{
			name:     "neither defined defaults to bare semver",
			metadata: metadataV1{},
			expected: []string{""},
		},
		{
			name: "only tag-prefix",
			metadata: metadataV1{
				Releases: struct {
					TagPrefix   string   `json:"tag-prefix"`
					TagPrefixes []string `json:"tag-prefixes"`
				}{TagPrefix: "my-sdk-"},
			},
			expected: []string{"my-sdk-"},
		},
		{
			name: "only tag-prefixes",
			metadata: metadataV1{
				Releases: struct {
					TagPrefix   string   `json:"tag-prefix"`
					TagPrefixes []string `json:"tag-prefixes"`
				}{TagPrefixes: []string{"", "my-sdk-"}},
			},
			expected: []string{"", "my-sdk-"},
		},
		{
			name: "both defined merges them",
			metadata: metadataV1{
				Releases: struct {
					TagPrefix   string   `json:"tag-prefix"`
					TagPrefixes []string `json:"tag-prefixes"`
				}{TagPrefix: "old-prefix-", TagPrefixes: []string{"", "new-prefix-"}},
			},
			expected: []string{"", "new-prefix-", "old-prefix-"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.metadata.effectivePrefixes()
			assert.Equal(t, tt.expected, got)
		})
	}
}
