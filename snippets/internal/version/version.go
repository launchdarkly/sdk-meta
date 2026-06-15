package version

// Version is the sdk-snippets release this binary was built from. Two
// independent paths set it:
//
//  1. release-please rewrites the literal below at release-PR time via
//     the snippets package's `extra-files` rule in release-please-
//     config.json. The `x-release-please-version` annotation must stay
//     on this exact line for that to work.
//  2. goreleaser passes -ldflags '-X .../version.Version={{.Version}}'
//     at build time, which the linker honours only for `var` (not
//     `const`) string declarations. This catches snapshot builds and
//     any future case where the source literal lags goreleaser.
//
// Either path alone is sufficient; using both is belt-and-braces.
var Version = "0.14.0" // x-release-please-version
