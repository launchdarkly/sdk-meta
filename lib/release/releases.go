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

type Raw struct {
	Tag  string `graphql:"tagName"`
	Date string `graphql:"publishedAt"`
}

type Release struct {
	Version string
	Date    time.Time
}

func (r Release) WithMajor() (WithMajor, error) {
	major, err := strconv.Atoi(strings.TrimPrefix(semver.Major(r.Version), "v"))
	if err != nil {
		return WithMajor{}, err
	}
	return WithMajor{r, major}, nil
}

type WithMajor struct {
	Release
	Major int
}

type WithEOL struct {
	Release
	EOL *time.Time
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

func (r WithEOL) MajorMinor() []string {
	return strings.Split(strings.TrimPrefix(semver.MajorMinor(r.Version), "v"), ".")
}

func (r WithEOL) MaybeEOL() *string {
	if r.EOL == nil {
		return nil
	}
	formatted := r.EOL.Format(time.RFC3339)
	return &formatted
}
func (r Release) Current() WithEOL {
	return WithEOL{r, nil}
}

func (r Release) WithEOL(t time.Time) WithEOL {
	return WithEOL{r, &t}
}

func (r Release) EOLPlusOneYear() time.Time {
	return r.Date.AddDate(1, 0, 0)
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

type Parser interface {
	Relevant(tag string) bool
	Parse(tag string) (string, error)
}

// basicParser parses tags of the form v[SEMVER] or [SEMVER].
type basicParser struct{}

func (p *basicParser) Relevant(tag string) bool {
	return semver.IsValid(tag) || semver.IsValid("v"+tag)
}

func (p *basicParser) Parse(tag string) (string, error) {
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

func (p *monorepoParser) Parse(tag string) (string, error) {
	return semver.Canonical(strings.TrimPrefix(tag, p.prefix)), nil
}

func Filter(releases []Raw, prefix string) ([]Release, error) {

	//const timeFormat = "2006-01-02T15:04:05Z"
	const timeFormat = time.RFC3339

	var parser Parser
	if prefix == "" {
		parser = &basicParser{}
	} else {
		parser = &monorepoParser{prefix: prefix}
	}

	var processed []Release
	for _, r := range releases {
		if !parser.Relevant(r.Tag) {
			continue
		}
		version, err := parser.Parse(r.Tag)
		if err != nil {
			return nil, err
		}
		date, err := time.Parse(timeFormat, r.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid release date for %s: %v", r.Tag, r.Date)
		}
		processed = append(processed, Release{Version: semver.Canonical(version), Date: date})
	}

	slices.SortFunc(processed, func(a Release, b Release) int {
		return -semver.Compare(a.Version, b.Version)
	})

	return processed, nil
}

func ExtractMajors(releases []Release) ([]WithMajor, error) {
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
