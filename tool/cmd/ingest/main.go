package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"time"

	"github.com/launchdarkly/sdk-meta/tool/lib/releases"
	_ "github.com/mattn/go-sqlite3"
	gh "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type metadataV1 struct {
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	UserAgents   []string `json:"userAgents"`
	WrapperNames []string `json:"wrapperNames"`
	Type         string   `json:"type"`
	Languages    []string `json:"languages"`
	Features     map[string]struct {
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
	schemaPath   string
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
	schema := flag.String("schema", "schemas/sdk_metadata.sql", "Path to schema file for database.")

	createDb := flag.Bool("create", false, "Create database if it does not exist")

	repo := flag.String("repo", "", "Github repo associated with the given metadata.json file in the form 'org/repo'")

	offline := flag.Bool("offline", false, "Don't fetch metadata that requires network access")

	flag.Parse()

	args := &args{
		*metadataPath,
		*dbPath,
		*schema,
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

func createOrOpen(path string, schema string, create bool) (*sql.DB, error) {
	if create {
		_ = os.Remove(path)
		if err := exec.Command("sh", "-c", fmt.Sprintf("sqlite3 %s < %s", path, schema)).Run(); err != nil {
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

	db, err := createOrOpen(args.dbPath, args.schemaPath, args.createDb)
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
		"repo": func(tx *sql.Tx, id string, metadata *metadataV1) error {
			if args.repo != "" {
				return insertRepo(tx, id, args.repo)
			}
			return nil
		},
		"features":     insertFeatures,
		"userAgents":   insertUserAgents,
		"wrapperNames": insertWrapperNames,
	}

	if !args.offline {
		if args.repo == "" {
			return fmt.Errorf("'repo' arg is required to run in online mode")
		}
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(context.Background(), src)

		client := gh.NewClient(httpClient)
		releaseCache := make(map[string][]releases.Raw)

		inserters["releases"] = func(tx *sql.Tx, sdkId string, metadata *metadataV1) error {
			if _, ok := releaseCache[args.repo]; !ok {
				rawReleases, err := releases.Query(client, args.repo)
				if err != nil {
					return err
				}
				releaseCache[args.repo] = rawReleases
			}
			all := releaseCache[args.repo]
			singleSDK, err := releases.Filter(all, metadata.Releases.TagPrefix)
			if err != nil {
				return err
			}

			return insertReleases(tx, sdkId, releases.Reduce(singleSDK))
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Sort SDK IDs for consistent iteration order
	sdkIds := make([]string, 0, len(metadata))
	for sdkId := range metadata {
		sdkIds = append(sdkIds, sdkId)
	}
	sort.Strings(sdkIds)

	for _, sdkId := range sdkIds {
		for column, insert := range inserters {
			if err := insert(tx, sdkId, metadata[sdkId]); err != nil {
				return fmt.Errorf("insert %s for %s: %v", column, sdkId, err)
			}
		}
	}

	return tx.Commit()
}

func insertLanguages(tx *sql.Tx, id string, metadata *metadataV1) error {
	if len(metadata.Languages) == 0 {
		return nil
	}
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
	if metadata.Type == "" {
		return nil
	}
	stmt, err := tx.Prepare("INSERT INTO sdk_types (id, type) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, metadata.Type)
	return err
}

func insertName(tx *sql.Tx, id string, metadata *metadataV1) error {
	if metadata.Name == "" {
		return nil
	}
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

func insertReleases(tx *sql.Tx, id string, release []releases.Parsed) error {
	stmt, err := tx.Prepare("INSERT INTO sdk_releases (id, major, minor, patch, date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, r := range release {
		v := r.Version
		_, err = stmt.Exec(id, v.Major(), v.Minor(), v.Patch(), r.Date.Format(time.RFC3339))
		if err != nil {
			return err
		}
	}
	return nil
}

func insertUserAgents(tx *sql.Tx, id string, metadata *metadataV1) error {
	if len(metadata.UserAgents) == 0 {
		return nil
	}
	stmt, err := tx.Prepare("INSERT INTO sdk_user_agents (id, userAgent) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, userAgent := range metadata.UserAgents {
		if _, err := stmt.Exec(id, userAgent); err != nil {
			return err
		}
	}

	return nil
}

func insertWrapperNames(tx *sql.Tx, id string, metadata *metadataV1) error {
	if len(metadata.WrapperNames) == 0 {
		return nil
	}
	stmt, err := tx.Prepare("INSERT INTO sdk_wrappers (id, wrapper) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, wrapper := range metadata.WrapperNames {
		if _, err := stmt.Exec(id, wrapper); err != nil {
			return err
		}
	}

	return nil
}
