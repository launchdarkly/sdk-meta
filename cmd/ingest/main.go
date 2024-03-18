package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/launchdarkly/sdk-meta/lib/eol"
	"github.com/launchdarkly/sdk-meta/lib/release"
	_ "github.com/mattn/go-sqlite3"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
	"time"
)

type metadataV1 struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	UserAgent string   `json:"user-agent"`
	Type      string   `json:"type"`
	Languages []string `json:"languages"`
	Features  map[string]struct {
		Introduced string  `json:"introduced"`
		Deprecated *string `json:"deprecated"`
		Removed    *string `json:"removed"`
	} `json:"features"`
	Releases struct {
		TagPrefix string `json:"tag-prefix"`
	} `json:"releases"`
}
type metadataCollection struct {
	Version int `json:"version"`
}

type args struct {
	metadataPath string
	dbPath       string
	createDb     bool
	repo         string
	offline      bool
}

func main() {
	metadataPath := flag.String("metadata", "metadata.json", "Path to metadata.json file")
	if *metadataPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	dbPath := flag.String("db", "metadata.sqlite3", "Path to database file. If not provided, a temp database will be used and discarded.")

	createDb := flag.Bool("create", false, "Create database if it does not exist")

	repo := flag.String("repo", "", "Github repo associated with the given metadata.json file in the form 'org/repo'")

	offline := flag.Bool("offline", false, "Don't fetch metadata that requires network access")

	flag.Parse()

	args := &args{
		*metadataPath,
		*dbPath,
		*createDb,
		*repo,
		*offline,
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

func createOrOpen(path string, create bool) (*sql.DB, error) {
	if create {
		_ = os.Remove(path)
		if err := exec.Command("sh", "-c", fmt.Sprintf("sqlite3 %s < ./schemas/sdk_metadata.sql", path)).Run(); err != nil {
			return nil, fmt.Errorf("couldn't create new database: %v", err)
		}
	}
	return sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=true&mode=rw&sync=full", path))
}

func run(args *args) error {
	metadata, err := parseMetadata(args.metadataPath)
	if err != nil {
		return err
	}

	db, err := createOrOpen(args.dbPath, args.createDb)
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

	if !args.offline {
		if args.repo == "" {
			return fmt.Errorf("'repo' arg is required to run in online mode")
		}
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(context.Background(), src)

		calculator := eol.NewCalculator(gh.NewClient(httpClient))

		inserters["releases"] = func(tx *sql.Tx, sdkId string, metadata *metadataV1) error {
			releases, err := calculator.Calculate(args.repo, metadata.Releases.TagPrefix)
			if err != nil {
				return err
			}
			return insertReleases(tx, sdkId, releases)
		}
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
	_, err = stmt.Exec(id, metadata.Name)
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

func insertReleases(tx *sql.Tx, id string, release []release.WithEOL) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_releases (id, major, minor, date, eol) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, release := range release {
		majorMinor := release.MajorMinor()
		_, err = stmt.Exec(id, majorMinor[0], majorMinor[1], release.Date.Format(time.RFC3339), release.MaybeEOL())
		if err != nil {
			return err
		}
	}
	return nil
}
