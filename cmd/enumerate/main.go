package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
	"slices"
)

type repo struct {
	name   string
	topics []string
}
type repositoriesQuery struct {
	Organization struct {
		Repositories struct {
			PageInfo struct {
				EndCursor   gh.String
				HasNextPage bool
			}
			Nodes []struct {
				Name             string
				IsPrivate        bool
				RepositoryTopics struct {
					Nodes []struct {
						Topic struct {
							Name string
						}
					}
				} `graphql:"repositoryTopics(first: 100)"`
			}
		} `graphql:"repositories(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $org)"`
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	repos, err := enumerate(gh.NewClient(httpClient))
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	repos = slices.DeleteFunc(repos, func(r repo) bool {
		return !slices.Contains(r.topics, "launchdarkly-sdk") || slices.Contains(r.topics, "sdk-examples") ||
			slices.Contains(r.topics, "launchdarkly-sdk-component")
	})

	var names []string
	for _, r := range repos {
		names = append(names, "launchdarkly/"+r.name)
	}

	if err := json.NewEncoder(os.Stdout).Encode(names); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

func enumerate(client *gh.Client) ([]repo, error) {
	variables := map[string]interface{}{
		"org":    gh.String("launchdarkly"),
		"cursor": (*gh.String)(nil),
	}

	var repos []repo

	var query repositoriesQuery
	for {
		err := client.Query(context.Background(), &query, variables)
		if err != nil {
			return nil, err
		}
		for _, n := range query.Organization.Repositories.Nodes {
			if n.IsPrivate {
				continue
			}
			var topics []string
			for _, t := range n.RepositoryTopics.Nodes {
				topics = append(topics, t.Topic.Name)
			}
			repos = append(repos, repo{
				name:   n.Name,
				topics: topics,
			})
		}
		if !query.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = gh.NewString(query.Organization.Repositories.PageInfo.EndCursor)
	}

	return repos, nil
}
