package validate

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestDockerBuildCommandWithoutCache(t *testing.T) {
	for _, tc := range []struct {
		name    string
		runner  *Runner
		builder string
	}{
		{"runner without build-cache", &Runner{}, "builder-1"},
		{"no builder in the environment", &Runner{BuildCache: "gha"}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := dockerBuildCommand(tc.runner, "haskell-server", "/v/Dockerfile", "img:abc", "/v", []string{"A=1"}, tc.builder)
			want := []string{"build", "--progress=plain", "--build-arg", "A=1", "-f", "/v/Dockerfile", "-t", "img:abc", "/v"}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("dockerBuildCommand() = %#v, want %#v", got, want)
			}
		})
	}
}

func TestDockerBuildCommandWithGHACache(t *testing.T) {
	got := dockerBuildCommand(&Runner{BuildCache: "gha"}, "haskell-server", "/v/Dockerfile", "img:abc", "/v", []string{"A=1"}, "builder-1")
	want := []string{
		"buildx", "build", "--builder", "builder-1", "--load", "--progress=plain",
		"--cache-from", "type=gha,scope=haskell-server",
		"--cache-to", "type=gha,mode=max,ignore-error=true,scope=haskell-server",
		"--build-arg", "A=1", "-f", "/v/Dockerfile", "-t", "img:abc", "/v",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("dockerBuildCommand() = %#v, want %#v", got, want)
	}
}

func TestLoadRunnerBuildCache(t *testing.T) {
	for _, tc := range []struct {
		name    string
		yaml    string
		wantErr string
	}{
		{"gha on docker", "mode: docker\nimage-prefix: x\nbuild-cache: gha\n", ""},
		{"unknown value", "mode: docker\nimage-prefix: x\nbuild-cache: s3\n", "build-cache must be `gha` or empty"},
		{"native mode", "mode: native\nbuild-cache: gha\n", "build-cache requires mode=docker"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			dir := filepath.Join(root, "languages", "test")
			if err := os.MkdirAll(dir, 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, "runner.yaml"), []byte(tc.yaml), 0o644); err != nil {
				t.Fatal(err)
			}
			r, _, err := loadRunner(root, "test")
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("loadRunner() error = %v", err)
				}
				if r.BuildCache != "gha" {
					t.Fatalf("BuildCache = %q, want gha", r.BuildCache)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("loadRunner() error = %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestBuildImageRejectsNativeRuntime(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "languages", "test")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "runner.yaml"), []byte("mode: native\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := BuildImage(root, "test", os.Stderr)
	if err == nil || !strings.Contains(err.Error(), "only docker validators have an image") {
		t.Fatalf("BuildImage() error = %v, want native-mode error", err)
	}
}
