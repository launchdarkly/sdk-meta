package snippets

import (
	"io/fs"
	"strings"
	"testing"
)

// The embedded FS is what the release binary ships with — consumers
// (gonfalon's snippets-sync action, the docs adapter) load from it
// without needing to check out sdk-meta. If this test starts failing,
// it almost certainly means the //go:embed directive lost track of
// the sdks/ tree and the binary would ship empty.
func TestSDKsFS_BundlesEverySDK(t *testing.T) {
	fsys := SDKsFS()

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		t.Fatalf("read embedded sdks/: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("embedded sdks/ is empty — go:embed isn't picking up the tree")
	}

	// Every SDK directory must carry an sdk.yaml; without it the
	// ld-application adapter's discoverTargetFiles would skip it.
	missing := []string{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if _, err := fs.Stat(fsys, e.Name()+"/sdk.yaml"); err != nil {
			missing = append(missing, e.Name())
		}
	}
	if len(missing) > 0 {
		t.Fatalf("SDKs missing sdk.yaml in embedded FS: %s", strings.Join(missing, ", "))
	}

	// Spot-check that at least one snippet file is present so we know
	// `all:` (not the default) is the embed mode — the default would
	// drop dotfiles and any underscore-prefixed paths, which we don't
	// have today but are easy to introduce later.
	count := 0
	_ = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && strings.HasSuffix(path, ".snippet.md") {
			count++
		}
		return nil
	})
	if count == 0 {
		t.Fatal("embedded FS contains no .snippet.md files")
	}
}
