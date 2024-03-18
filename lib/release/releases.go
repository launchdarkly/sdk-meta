package release

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

// How long we support the latest SDK version.
const supportWindowYears = 1

// Raw is the raw tag data returned from the github GraphQL releases query.
type Raw struct {
	Tag  string `graphql:"tagName"`
	Date string `graphql:"publishedAt"`
}

// Parsed is the post-processed version of a Raw structure, with the version extracted and date
// parsed.
type Parsed struct {
	Version string
	Date    time.Time
}

// WithMajor annotates a parsed release with a comparable major version.
type WithMajor struct {
	Parsed
	Major int
}

// WithEOL annotates a parsed release with an optional EOL date. If nil, this release has no EOL.
type WithEOL struct {
	Parsed
	EOL *time.Time
}

// WithMajor extracts the major version of a release and returns a WithMajor containing it.
func (r Parsed) WithMajor() (WithMajor, error) {
	major, err := strconv.Atoi(strings.TrimPrefix(semver.Major(r.Version), "v"))
	if err != nil {
		return WithMajor{}, err
	}
	return WithMajor{r, major}, nil
}

// AsCurrent marks this release as current.
func (r Parsed) AsCurrent() WithEOL {
	return WithEOL{r, nil}
}

// AsExpiring marks this release as eventually going EOL on a timestamp.
func (r Parsed) AsExpiring(t time.Time) WithEOL {
	return WithEOL{r, &t}
}

// SupportWindow returns the point in time where this release would expire, if it were current.
func (r Parsed) SupportWindow() time.Time {
	return r.Date.AddDate(supportWindowYears, 0, 0)
}

// MajorMinor returns a slice containing the major and minor version, e.g. ["1" (major), "2" (minor)]
func (r WithEOL) MajorMinor() []string {
	return strings.Split(strings.TrimPrefix(semver.MajorMinor(r.Version), "v"), ".")
}

// MaybeEOL returns an RFC3339 timestamp of the EOL date of this release, or nil if there is no EOL.
func (r WithEOL) MaybeEOL() *string {
	if r.EOL == nil {
		return nil
	}
	formatted := r.EOL.Format(time.RFC3339)
	return &formatted
}

type releasesQuery struct {
	Repository struct {
		Releases struct {
			Nodes    []Raw
			PageInfo struct {
				EndCursor   gh.String
				HasNextPage bool
			}
		} `graphql:"releases(first: 100, after: $cursor)"`
	} `graphql:"repository(owner: $org, name: $repo)"`
}

func Query(client *gh.Client,
	repoPath string) ([]Raw, error) {
	parts := strings.Split(repoPath, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repo path: %s", repoPath)
	}

	org := parts[0]
	repo := parts[1]

	variables := map[string]interface{}{
		"org":    gh.String(org),
		"repo":   gh.String(repo),
		"cursor": (*gh.String)(nil),
	}

	var releases []Raw

	var query releasesQuery
	for {
		err := client.Query(context.Background(), &query, variables)
		if err != nil {
			return nil, err
		}
		releases = append(releases, query.Repository.Releases.Nodes...)
		if !query.Repository.Releases.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = gh.NewString(query.Repository.Releases.PageInfo.EndCursor)
	}

	return releases, nil
}

type TagParser interface {
	// Relevant returns true if the given tag is relevant to the parser, or should be skipped.
	Relevant(tag string) bool
	// ParseSemver returns the semantic version associated with the tag, or an error. The semver should contain
	// a leading 'v'.
	ParseSemver(tag string) (string, error)
}

// basicParser parses tags of the form v[SEMVER] or [SEMVER].
type basicParser struct{}

func (p *basicParser) Relevant(tag string) bool {
	return semver.IsValid(tag) || semver.IsValid("v"+tag)
}

func (p *basicParser) ParseSemver(tag string) (string, error) {
	if strings.HasPrefix(tag, "v") {
		return semver.Canonical(tag), nil
	}
	return semver.Canonical("v" + tag), nil
}

// monorepoParser parses tags of the form [PREFIX][SEMVER].
type monorepoParser struct {
	prefix string
}

func (p *monorepoParser) Relevant(tag string) bool {
	return strings.HasPrefix(tag, p.prefix) && semver.IsValid(strings.TrimPrefix(tag, p.prefix))
}

func (p *monorepoParser) ParseSemver(tag string) (string, error) {
	return semver.Canonical(strings.TrimPrefix(tag, p.prefix)), nil
}

func Filter(releases []Raw, prefix string) ([]Parsed, error) {

	const timeFormat = time.RFC3339

	var parser TagParser
	if prefix == "" {
		parser = &basicParser{}
	} else {
		parser = &monorepoParser{prefix: prefix}
	}

	var processed []Parsed
	for _, r := range releases {
		if !parser.Relevant(r.Tag) {
			continue
		}
		version, err := parser.ParseSemver(r.Tag)
		if err != nil {
			return nil, err
		}
		date, err := time.Parse(timeFormat, r.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid release date for %s: %v", r.Tag, r.Date)
		}
		processed = append(processed, Parsed{Version: semver.Canonical(version), Date: date})
	}

	slices.SortFunc(processed, func(a Parsed, b Parsed) int {
		return -semver.Compare(a.Version, b.Version)
	})

	return processed, nil
}

func ExtractMajors(releases []Parsed) ([]WithMajor, error) {
	var withMajors []WithMajor
	for _, r := range releases {
		withMajor, err := r.WithMajor()
		if err != nil {
			return nil, err
		}
		withMajors = append(withMajors, withMajor)
	}
	return withMajors, nil
}
