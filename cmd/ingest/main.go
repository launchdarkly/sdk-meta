package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
	"strings"
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

func releases(client *gh.Client, org string, repo string) (*releasesQuery, error) {
	var releasesQuery releasesQuery
	err := client.Query(context.Background(), &releasesQuery, map[string]interface{}{
		"org":  gh.String(org),
		"repo": gh.String(repo),
	})
	if err != nil {
		return nil, err
	}
	return &releasesQuery, nil
}

func run2() error {
	// Parse arg for -repo using flag package.

	repoPath := flag.String("repo", "", "The repository to crawl (org/repo syntax)")
	flag.Parse()
	if *repoPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	parts := strings.Split(*repoPath, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid repo path: %s", *repoPath)
	}

	org := parts[0]
	repo := parts[1]

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := gh.NewClient(httpClient)

	rels, err := releases(client, org, repo)
	if err != nil {
		return err
	}
	fmt.Println(rels)
	return nil
}

type metadataV1 struct {
	Path      string   `json:"path"`
	UserAgent string   `json:"user-agent"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Features  map[string]struct {
		Introduced string  `json:"introduced"`
		Deprecated *string `json:"deprecated"`
		Removed    *string `json:"removed"`
	} `json:"features"`
}
type metadataCollection struct {
	Version int `json:"version"`
}

type args struct {
	metadataPath string
	dbPath       string
	repo         string
}

func main() {
	metadataPath := flag.String("metadata", "metadata.json", "Path to metadata.json file")
	if *metadataPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	dbPath := flag.String("db", "sdk_metadata.sqlite3", "Path to database file")

	repo := flag.String("repo", "", "Github repo associated with the given metadata.json file in the form 'org/repo'")
	flag.Parse()

	args := &args{
		*metadataPath,
		*dbPath,
		*repo,
	}

	if err := run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseMetadata(path string) (map[string]*metadataV1, error) {
	metadataFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()

	var collection metadataCollection
	if err := json.NewDecoder(metadataFile).Decode(&collection); err != nil {
		return nil, err
	}

	if _, err := metadataFile.Seek(0, 0); err != nil {
		return nil, err
	}

	switch collection.Version {
	case 1:
		var container struct {
			Sdks map[string]*metadataV1 `json:"sdks"`
		}
		if err := json.NewDecoder(metadataFile).Decode(&container); err != nil {
			return nil, err
		}
		return container.Sdks, nil
	default:
		return nil, fmt.Errorf("unknown metadata version: %d", collection.Version)
	}
}

func run(args *args) error {
	metadata, err := parseMetadata(args.metadataPath)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=true&mode=rw&sync=full", args.dbPath))
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	inserters := map[string]func(*sql.Tx, string, *metadataV1) error{
		"languages": insertLanguages,
		"type":      insertType,
		"name":      insertName,
		"repo": func(tx *sql.Tx, sdkId string, metadata *metadataV1) error {
			if args.repo != "" {
				return insertRepo(tx, sdkId, args.repo)
			}
			return nil
		},
		"features": insertFeatures,
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for sdkId, metadata := range metadata {
		for column, insert := range inserters {
			if err := insert(tx, sdkId, metadata); err != nil {
				return fmt.Errorf("insert %s for %s: %v", column, sdkId, err)
			}
		}
	}

	return tx.Commit()
}

func insertLanguages(tx *sql.Tx, id string, metadata *metadataV1) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_languages (id, language) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, language := range metadata.Languages {
		if _, err := stmt.Exec(id, language); err != nil {
			return err
		}
	}

	return nil
}

func insertType(tx *sql.Tx, id string, metadata *metadataV1) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_types (id, type) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, metadata.Type)
	return err
}

func insertName(tx *sql.Tx, id string, metadata *metadataV1) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_names (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, metadata.Path)
	return err
}

func insertRepo(tx *sql.Tx, id string, repo string) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_repos (id, github) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, repo)
	return err
}

func insertFeatures(tx *sql.Tx, id string, metadata *metadataV1) error {
	// Todo: how to handle the empty string/nil values
	stmt, err := tx.Prepare("INSERT INTO sdk_features (id, feature, introduced, deprecated, removed) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for feature, info := range metadata.Features {
		_, err = stmt.Exec(id, feature, info.Introduced, info.Deprecated, info.Removed)
		if err != nil {
			return err
		}
	}
	return nil
}
