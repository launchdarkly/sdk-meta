package validate

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

// sdkVersionSpec is a floating SDK version entry from shared/versions/sdks.env,
// written as <registry>:<package>@<family>. Family is a numeric version
// prefix: "4" selects the newest stable 4.x.y, "0.2" the newest 0.2.x.
type sdkVersionSpec struct {
	Registry string
	Package  string
	Family   []int
	raw      string
}

func (s sdkVersionSpec) String() string { return s.raw }

var sdkRegistries = map[string]func(pkg string) ([]sdkVersionCandidate, error){
	"npm":        npmVersions,
	"pub":        pubVersions,
	"maven":      mavenVersions,
	"hackage":    hackageVersions,
	"hex":        hexVersions,
	"luarocks":   luarocksVersions,
	"github-tag": githubTagVersions,
}

func parseSDKVersionSpec(raw string) (sdkVersionSpec, error) {
	registry, rest, ok := strings.Cut(raw, ":")
	at := strings.LastIndex(rest, "@")
	if !ok || at <= 0 || at == len(rest)-1 {
		return sdkVersionSpec{}, fmt.Errorf("invalid SDK version spec %q: want <registry>:<package>@<family>", raw)
	}
	if _, ok := sdkRegistries[registry]; !ok {
		return sdkVersionSpec{}, fmt.Errorf("invalid SDK version spec %q: unknown registry %q", raw, registry)
	}
	family, ok := parseNumericVersion(rest[at+1:])
	if !ok {
		return sdkVersionSpec{}, fmt.Errorf("invalid SDK version spec %q: family must be numeric, such as 4 or 0.2", raw)
	}
	pkg := rest[:at]
	if registry == "maven" || registry == "luarocks" || registry == "github-tag" {
		sep := map[string]string{"maven": ":", "luarocks": "/", "github-tag": ":"}[registry]
		if a, b, ok := strings.Cut(pkg, sep); !ok || a == "" || (b == "" && registry != "github-tag") {
			return sdkVersionSpec{}, fmt.Errorf("invalid SDK version spec %q: malformed %s package %q", raw, registry, pkg)
		}
	}
	return sdkVersionSpec{Registry: registry, Package: pkg, Family: family, raw: raw}, nil
}

// sdkVersionCandidate is one published release: Value is what gets passed to
// the Dockerfile (a version, a rock version, or a full git tag) and Version
// is its numeric form used for family matching and ordering.
type sdkVersionCandidate struct {
	Value   string
	Version []int
}

// parseNumericVersion accepts dotted non-negative integers only, so
// prereleases and build metadata never satisfy a family.
func parseNumericVersion(v string) ([]int, bool) {
	if v == "" {
		return nil, false
	}
	parts := strings.Split(v, ".")
	out := make([]int, len(parts))
	for i, p := range parts {
		if p == "" || strings.TrimLeft(p, "0123456789") != "" {
			return nil, false
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, false
		}
		out[i] = n
	}
	return out, true
}

func stableCandidates(versions []string) []sdkVersionCandidate {
	var out []sdkVersionCandidate
	for _, v := range versions {
		if parsed, ok := parseNumericVersion(v); ok {
			out = append(out, sdkVersionCandidate{Value: v, Version: parsed})
		}
	}
	return out
}

func compareVersions(a, b []int) int {
	for i := 0; i < len(a) || i < len(b); i++ {
		var x, y int
		if i < len(a) {
			x = a[i]
		}
		if i < len(b) {
			y = b[i]
		}
		if x != y {
			if x < y {
				return -1
			}
			return 1
		}
	}
	return len(a) - len(b)
}

func newestInFamily(family []int, candidates []sdkVersionCandidate) (string, bool) {
	var best *sdkVersionCandidate
	for i := range candidates {
		c := &candidates[i]
		if len(c.Version) <= len(family) {
			continue
		}
		match := true
		for j, n := range family {
			if c.Version[j] != n {
				match = false
				break
			}
		}
		if match && (best == nil || compareVersions(c.Version, best.Version) > 0) {
			best = c
		}
	}
	if best == nil {
		return "", false
	}
	return best.Value, true
}

type sdkResolution struct {
	once  sync.Once
	value string
	err   error
}

var sdkResolutions sync.Map

// resolveSDKVersion returns the newest stable release in spec's family,
// looking each spec up at most once per process. The first caller logs the
// resolution to out so every validator run records what it tested.
func resolveSDKVersion(spec sdkVersionSpec, out io.Writer) (string, error) {
	entry, _ := sdkResolutions.LoadOrStore(spec.raw, &sdkResolution{})
	r := entry.(*sdkResolution)
	r.once.Do(func() {
		candidates, err := sdkRegistries[spec.Registry](spec.Package)
		if err != nil {
			r.err = fmt.Errorf("resolve %s: %w", spec, err)
			return
		}
		value, ok := newestInFamily(spec.Family, candidates)
		if !ok {
			r.err = fmt.Errorf("resolve %s: no stable release in that family", spec)
			return
		}
		r.value = value
		fmt.Fprintf(out, "resolved %s -> %s\n", spec, value)
	})
	return r.value, r.err
}

var (
	npmRegistryURL      = "https://registry.npmjs.org"
	pubRegistryURL      = "https://pub.dev"
	mavenRepositoryURLs = []string{
		"https://repo1.maven.org/maven2",
		"https://maven-central.storage-download.googleapis.com/maven2",
	}
	hackageURL  = "https://hackage.haskell.org"
	hexURL      = "https://hex.pm"
	luarocksURL = "https://luarocks.org"
	gitLsRemote = func(repoURL, pattern string) ([]byte, error) {
		return exec.Command("git", "ls-remote", "--tags", "--refs", repoURL, pattern).Output()
	}
	registryClient = &http.Client{Timeout: 30 * time.Second}
)

func fetchRegistry(rawURL, accept string) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt < 4; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt*attempt) * time.Second)
		}
		req, err := http.NewRequest(http.MethodGet, rawURL, nil)
		if err != nil {
			return nil, err
		}
		if accept != "" {
			req.Header.Set("Accept", accept)
		}
		resp, err := registryClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}
		if resp.StatusCode == http.StatusOK {
			return body, nil
		}
		lastErr = fmt.Errorf("GET %s: HTTP %d", rawURL, resp.StatusCode)
		if resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode < 500 {
			break
		}
	}
	return nil, lastErr
}

func npmVersions(pkg string) ([]sdkVersionCandidate, error) {
	body, err := fetchRegistry(npmRegistryURL+"/"+url.PathEscape(pkg), "application/vnd.npm.install-v1+json")
	if err != nil {
		return nil, err
	}
	var doc struct {
		Versions map[string]json.RawMessage `json:"versions"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse npm metadata for %s: %w", pkg, err)
	}
	versions := make([]string, 0, len(doc.Versions))
	for v := range doc.Versions {
		versions = append(versions, v)
	}
	return stableCandidates(versions), nil
}

func pubVersions(pkg string) ([]sdkVersionCandidate, error) {
	body, err := fetchRegistry(pubRegistryURL+"/api/packages/"+url.PathEscape(pkg), "application/vnd.pub.v2+json")
	if err != nil {
		return nil, err
	}
	var doc struct {
		Versions []struct {
			Version   string `json:"version"`
			Retracted bool   `json:"retracted"`
		} `json:"versions"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse pub.dev metadata for %s: %w", pkg, err)
	}
	var versions []string
	for _, v := range doc.Versions {
		if !v.Retracted {
			versions = append(versions, v.Version)
		}
	}
	return stableCandidates(versions), nil
}

func mavenVersions(pkg string) ([]sdkVersionCandidate, error) {
	group, artifact, _ := strings.Cut(pkg, ":")
	path := strings.ReplaceAll(group, ".", "/") + "/" + artifact + "/maven-metadata.xml"
	var errs []string
	for _, base := range mavenRepositoryURLs {
		body, err := fetchRegistry(base+"/"+path, "")
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		var doc struct {
			Versions []string `xml:"versioning>versions>version"`
		}
		if err := xml.Unmarshal(body, &doc); err != nil {
			errs = append(errs, fmt.Sprintf("parse %s/%s: %v", base, path, err))
			continue
		}
		return stableCandidates(doc.Versions), nil
	}
	return nil, fmt.Errorf("%s", strings.Join(errs, "; "))
}

func hackageVersions(pkg string) ([]sdkVersionCandidate, error) {
	body, err := fetchRegistry(hackageURL+"/package/"+url.PathEscape(pkg)+"/preferred", "application/json")
	if err != nil {
		return nil, err
	}
	var doc struct {
		Normal []string `json:"normal-version"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse Hackage metadata for %s: %w", pkg, err)
	}
	return stableCandidates(doc.Normal), nil
}

func hexVersions(pkg string) ([]sdkVersionCandidate, error) {
	body, err := fetchRegistry(hexURL+"/api/packages/"+url.PathEscape(pkg), "application/json")
	if err != nil {
		return nil, err
	}
	var doc struct {
		Releases []struct {
			Version string `json:"version"`
		} `json:"releases"`
		Retirements map[string]json.RawMessage `json:"retirements"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse Hex metadata for %s: %w", pkg, err)
	}
	var versions []string
	for _, r := range doc.Releases {
		if _, retired := doc.Retirements[r.Version]; !retired {
			versions = append(versions, r.Version)
		}
	}
	return stableCandidates(versions), nil
}

// luarocksVersions reads a publisher's manifest; pkg is <user>/<rock>. Rock
// versions carry a rockspec revision ("2.2.0-0") that is kept in Value and
// ordered as a trailing component.
func luarocksVersions(pkg string) ([]sdkVersionCandidate, error) {
	user, rock, _ := strings.Cut(pkg, "/")
	body, err := fetchRegistry(luarocksURL+"/manifests/"+url.PathEscape(user)+"/manifest-5.3.json", "application/json")
	if err != nil {
		return nil, err
	}
	var doc struct {
		Repository map[string]map[string]json.RawMessage `json:"repository"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse LuaRocks manifest for %s: %w", user, err)
	}
	var out []sdkVersionCandidate
	for v := range doc.Repository[rock] {
		base, rev, ok := strings.Cut(v, "-")
		parsed, baseOK := parseNumericVersion(base)
		revision, revErr := strconv.Atoi(rev)
		if !ok || !baseOK || revErr != nil {
			continue
		}
		out = append(out, sdkVersionCandidate{Value: v, Version: append(parsed, revision)})
	}
	return out, nil
}

// githubTagVersions lists tags of a GitHub repository; pkg is
// <owner>/<repo>:<tag-prefix>, and Value is the full tag name.
func githubTagVersions(pkg string) ([]sdkVersionCandidate, error) {
	repo, prefix, _ := strings.Cut(pkg, ":")
	var output []byte
	var err error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt*attempt) * time.Second)
		}
		if output, err = gitLsRemote("https://github.com/"+repo, "refs/tags/"+prefix+"*"); err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("git ls-remote %s: %w", repo, err)
	}
	var out []sdkVersionCandidate
	for _, line := range strings.Split(string(output), "\n") {
		_, ref, ok := strings.Cut(strings.TrimSpace(line), "\t")
		if !ok {
			continue
		}
		tag := strings.TrimPrefix(ref, "refs/tags/")
		if !strings.HasPrefix(tag, prefix) {
			continue
		}
		if parsed, ok := parseNumericVersion(strings.TrimPrefix(tag, prefix)); ok {
			out = append(out, sdkVersionCandidate{Value: tag, Version: parsed})
		}
	}
	return out, nil
}
