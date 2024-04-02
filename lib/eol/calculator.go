package eol

import (
	"github.com/launchdarkly/sdk-meta/lib/releases"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/mod/semver"
	"slices"
)

type Calculator struct {
	client *gh.Client
	cache  map[string][]releases.Raw
}

func NewCalculator(client *gh.Client) *Calculator {
	return &Calculator{
		client: client,
		cache:  make(map[string][]releases.Raw),
	}
}

func (e *Calculator) Calculate(repo string, prefix string) ([]releases.WithEOL, error) {
	_, ok := e.cache[repo]
	if !ok {
		rawReleases, err := releases.Query(e.client, repo)
		if err != nil {
			return nil, err
		}
		e.cache[repo] = rawReleases
	}
	filteredReleases, err := releases.Filter(e.cache[repo], prefix)
	if err != nil {
		return nil, err
	}
	releasesWithMajors, err := releases.ExtractMajors(filteredReleases)
	if err != nil {
		return nil, err
	}
	return CalculateEOLs(releasesWithMajors), nil
}

func CalculateEOLs(releases []releases.WithMajor) []releases.WithEOL {
	// First, delete irrelvant releases (those before major version 1)

	releases = slices.DeleteFunc(releases, func(a releases.WithMajor) bool {
		return a.Major == 0
	})

	// Second, sort ascending so that older releases come first in the array.
	// This is necessary so that the CompactFunc keeps the *oldest* which is relevant for the EOL
	// calculation.
	slices.SortFunc(releases, func(a releases.WithMajor, b releases.WithMajor) int {
		return semver.Compare(a.Version, b.Version)
	})

	releases = slices.CompactFunc(releases, func(a releases.WithMajor, b releases.WithMajor) bool {
		return semver.MajorMinor(a.Version) == semver.MajorMinor(b.Version)
	})

	// Before running the EOL algorithm, reverse it so that the first entry is the latest release.
	slices.Reverse(releases)

	var releasesWithEOL []releases.WithEOL
	for i := range releases {
		releasesWithEOL = append(releasesWithEOL, eol(i, releases))
	}
	return releasesWithEOL
}

func eol(i int, releases []releases.WithMajor) releases.WithEOL {
	this := releases[i]
	if i == 0 {
		return this.AsCurrent()
	}

	nextMajor := getNextMajor(releases, i)
	nextRelease := i - 1

	// If this is a version within the currently supported major version,
	// then the EOL date is the next release + 1 year.
	if nextMajor == i {
		return this.AsExpiring(releases[nextRelease].SupportWindow())
	}

	// Otherwise, there's another major version out: cap the EOL date
	// at min(nextMajor + 1 year, nextRelease + 1 year). Otherwise, releases
	// to old major branches will indefinitely extend the support window.

	nextMajorEOL := releases[nextMajor].SupportWindow()
	nextReleaseEOL := releases[nextRelease].SupportWindow()

	if nextMajorEOL.Compare(nextReleaseEOL) == -1 {
		return this.AsExpiring(nextMajorEOL)
	}
	return this.AsExpiring(nextReleaseEOL)
}

func getNextMajor(releases []releases.WithMajor, i int) int {
	currentMajor := releases[i].Major

	for j := i - 1; j >= 0; j-- {
		if releases[j].Major > currentMajor {
			return j
		}
	}
	return i
}
