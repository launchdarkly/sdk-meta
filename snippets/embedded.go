// Package snippets exposes the canonical SDK snippet sources embedded
// at build time. Consumer-side tooling (gonfalon's snippets-sync action,
// the future ld-docs adapter, the validator dispatcher) loads from this
// embedded FS by default — so a release artifact ships the engine and
// the snippets together, atomically. Local development overrides the
// default via `--sdks=./sdks`, which is just `os.DirFS(path)` instead.
//
// The directive only covers `sdks/`. Validator harnesses (Dockerfiles,
// shell scripts) stay on disk in the repo because they're sdk-meta
// internal CI — consumers never run `validate`.
package snippets

import (
	"embed"
	"io/fs"
)

//go:embed all:sdks
var bundled embed.FS

// SDKsFS returns the embedded sdks/ tree as an fs.FS. The returned FS is
// rooted at the sdks/ directory's contents (so `<sdk-id>/sdk.yaml` is the
// path of an sdk descriptor), matching `os.DirFS("./sdks")`.
func SDKsFS() fs.FS {
	sub, err := fs.Sub(bundled, "sdks")
	if err != nil {
		panic(err)
	}
	return sub
}
