package validate

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeValidatorVersionFiles(t *testing.T, root string, images, npm, toolchains string) {
	t.Helper()
	dir := filepath.Join(root, "shared", "versions")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"images.env":     images,
		"npm.env":        npm,
		"toolchains.env": toolchains,
	}
	for name, contents := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(contents), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestLoadValidatorVersions(t *testing.T) {
	root := t.TempDir()
	writeValidatorVersionFiles(t, root,
		"# comment\n\nIMAGE=ubuntu:24.04\n",
		"# comment\nNPM_VERSION=12\n",
		"GO_VERSION=1.25.12\n",
	)

	got, err := loadValidatorVersions(root)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"IMAGE":       "ubuntu:24.04",
		"NPM_VERSION": "12",
		"GO_VERSION":  "1.25.12",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("loadValidatorVersions() = %#v, want %#v", got, want)
	}
}

func TestLoadValidatorVersionsRejectsMalformedEntry(t *testing.T) {
	root := t.TempDir()
	writeValidatorVersionFiles(t, root,
		"IMAGE=ubuntu:24.04\nnot an assignment\n",
		"NPM_VERSION=12\n",
		"GO_VERSION=1.25.12\n",
	)

	_, err := loadValidatorVersions(root)
	if err == nil || !strings.Contains(err.Error(), "invalid entry") {
		t.Fatalf("loadValidatorVersions() error = %v, want malformed-entry error", err)
	}
}

func TestLoadValidatorVersionsRejectsDuplicateKeyAcrossFiles(t *testing.T) {
	root := t.TempDir()
	writeValidatorVersionFiles(t, root,
		"SHARED=value\n",
		"SHARED=other\n",
		"GO_VERSION=1.25.12\n",
	)

	_, err := loadValidatorVersions(root)
	if err == nil || !strings.Contains(err.Error(), `duplicate validator version key "SHARED"`) {
		t.Fatalf("loadValidatorVersions() error = %v, want duplicate-key error", err)
	}
}

func TestDockerfileArgs(t *testing.T) {
	dockerfile := filepath.Join(t.TempDir(), "Dockerfile")
	contents := "ARG FIRST=one\nFROM ${FIRST}\nARG SECOND\nARG FIRST\nARG THIRD = three\n"
	if err := os.WriteFile(dockerfile, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := dockerfileArgs(dockerfile)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"FIRST", "SECOND", "THIRD"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("dockerfileArgs() = %#v, want %#v", got, want)
	}
}

func TestBuildImageFailsForMissingDeclaredArg(t *testing.T) {
	root := t.TempDir()
	writeValidatorVersionFiles(t, root,
		"KNOWN=value\n",
		"NPM_VERSION=12\n",
		"GO_VERSION=1.25.12\n",
	)
	runnerDir := filepath.Join(root, "languages", "test")
	if err := os.MkdirAll(runnerDir, 0o755); err != nil {
		t.Fatal(err)
	}
	dockerfile := filepath.Join(runnerDir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("ARG MISSING\nFROM ${MISSING}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	err := buildImage(Config{ValidatorsDir: root}, &Runner{ImagePrefix: "test-validator"}, runnerDir, &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "declares ARG MISSING but no version entry exists") {
		t.Fatalf("buildImage() error = %v, want missing-ARG error", err)
	}
}
