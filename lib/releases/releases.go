package releases

import (
	"context"
	"fmt"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/mod/semver"
	"slices"
	"strconv"
	"strings"
	"time"
)

type releasesQuery struct {
	Repository struct {
		Releases struct {
			Nodes []struct {
				TagName     string
				PublishedAt string
			}
		} `graphql:"releases(last: 100)"`
	} `graphql:"repository(owner: $org, name: $repo)"`
}

type ReleaseWithEOL struct {
	Release
	EOL *time.Time
}

type Release struct {
	Version string
	Major   int
	Date    time.Time
}

func (r ReleaseWithEOL) MajorMinor() string {
	return strings.TrimPrefix(semver.MajorMinor(r.Version), "v")
}

func (r ReleaseWithEOL) MaybeEOL() *string {
	if r.EOL == nil {
		return nil
	}
	formatted := r.EOL.Format(time.RFC3339)
	return &formatted
}
func (r Release) current() ReleaseWithEOL {
	return ReleaseWithEOL{r, nil}
}

func (r Release) withEOL(t time.Time) ReleaseWithEOL {
	return ReleaseWithEOL{r, &t}
}

func (r Release) eolPlusOneYear() time.Time {
	return r.Date.AddDate(1, 0, 0)
}

type Fetcher struct {
	client *gh.Client
	cache  map[string]*releasesQuery
}

func New(client *gh.Client) *Fetcher {
	return &Fetcher{
		client: client,
		cache:  make(map[string]*releasesQuery),
	}
}

func (f *Fetcher) Fetch(repoPath string, tagPrefix string) ([]Release, error) {
	parts := strings.Split(repoPath, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repo path: %s", repoPath)
	}

	org := parts[0]
	repo := parts[1]

	_, ok := f.cache[repoPath]
	if !ok {
		var releasesQuery releasesQuery
		err := f.client.Query(context.Background(), &releasesQuery, map[string]interface{}{
			"org":  gh.String(org),
			"repo": gh.String(repo),
		})
		if err != nil {
			return nil, err
		}
		f.cache[repoPath] = &releasesQuery
	}

	var releases []Release

	const timeFormat = "2006-01-02T15:04:05Z"

	for _, node := range f.cache[repoPath].Repository.Releases.Nodes {
		withoutPrefix := strings.TrimPrefix(node.TagName, tagPrefix)
		if semver.IsValid(withoutPrefix) {
			parsedDate, err := time.Parse(timeFormat, node.PublishedAt)
			if err != nil {
				return nil, fmt.Errorf("invalid release encountered in %s (%s): %s", repoPath, node.TagName, node.PublishedAt)
			}

			canonical := semver.Canonical(withoutPrefix)

			release := Release{
				Version: canonical,
				Date:    parsedDate,
			}

			major, err := strconv.Atoi(strings.TrimPrefix(semver.Major(canonical), "v"))
			if err != nil {
				return nil, fmt.Errorf("invalid major version in %s (%s): %s", repoPath, node.TagName, release.Version)
			}

			release.Major = major

			releases = append(releases, release)
		}
	}

	slices.SortFunc(releases, func(a Release, b Release) int {
		return -semver.Compare(a.Version, b.Version)
	})

	return releases, nil
}

func CalculateEOLs(releases []Release) []ReleaseWithEOL {
	releases = slices.CompactFunc(releases, func(a Release, b Release) bool {
		return semver.MajorMinor(a.Version) == semver.MajorMinor(b.Version)
	})

	releases = slices.DeleteFunc(releases, func(a Release) bool {
		return a.Major == 0
	})

	var releasesWithEOL []ReleaseWithEOL
	for i := range releases {
		releasesWithEOL = append(releasesWithEOL, eol(i, releases))
	}
	return releasesWithEOL
}

func eol(i int, releases []Release) ReleaseWithEOL {
	this := releases[i]
	if i == 0 {
		return this.current()
	}

	nextMajor := getNextMajor(releases, i)
	nextRelease := i - 1

	// If this is a version within the currently supported major version,
	// then the EOL date is the next release + 1 year.
	if nextMajor == i {
		return this.withEOL(releases[nextRelease].eolPlusOneYear())
	}

	// Otherwise, there's another major version out: cap the EOL date
	// at min(nextMajor + 1 year, nextRelease + 1 year). Otherwise, releases
	// to old major branches will indefinitely extend the support window.

	nextMajorEOL := releases[nextMajor].eolPlusOneYear()
	nextReleaseEOL := releases[nextRelease].eolPlusOneYear()

	if nextMajorEOL.Compare(nextReleaseEOL) == -1 {
		return this.withEOL(nextMajorEOL)
	}
	return this.withEOL(nextReleaseEOL)
}

func getNextMajor(releases []Release, i int) int {
	currentMajor := releases[i].Major

	for j := i - 1; j >= 0; j-- {
		if releases[j].Major > currentMajor {
			return j
		}
	}
	return i
}
