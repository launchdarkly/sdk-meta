package validate

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseSDKVersionSpec(t *testing.T) {
	cases := []struct {
		raw      string
		registry string
		pkg      string
		family   []int
	}{
		{"npm:@launchdarkly/js-client-sdk@4", "npm", "@launchdarkly/js-client-sdk", []int{4}},
		{"npm:@launchdarkly/fastly-server-sdk@0.2", "npm", "@launchdarkly/fastly-server-sdk", []int{0, 2}},
		{"maven:com.launchdarkly:launchdarkly-android-client-sdk@5", "maven", "com.launchdarkly:launchdarkly-android-client-sdk", []int{5}},
		{"luarocks:launchdarkly/launchdarkly-server-sdk@2", "luarocks", "launchdarkly/launchdarkly-server-sdk", []int{2}},
		{"github-tag:launchdarkly/cpp-sdks:launchdarkly-cpp-server-v@3", "github-tag", "launchdarkly/cpp-sdks:launchdarkly-cpp-server-v", []int{3}},
	}
	for _, c := range cases {
		got, err := parseSDKVersionSpec(c.raw)
		if err != nil {
			t.Fatalf("parseSDKVersionSpec(%q): %v", c.raw, err)
		}
		if got.Registry != c.registry || got.Package != c.pkg || !reflect.DeepEqual(got.Family, c.family) {
			t.Fatalf("parseSDKVersionSpec(%q) = %#v", c.raw, got)
		}
	}
	for _, raw := range []string{
		"4.10.3",
		"npm:@launchdarkly/js-client-sdk",
		"npm:@launchdarkly/js-client-sdk@^4",
		"npm:@launchdarkly/js-client-sdk@4.x",
		"cargo:launchdarkly-server-sdk@2",
		"maven:launchdarkly-android-client-sdk@5",
		"luarocks:launchdarkly-server-sdk@2",
	} {
		if _, err := parseSDKVersionSpec(raw); err == nil {
			t.Fatalf("parseSDKVersionSpec(%q) succeeded, want error", raw)
		}
	}
}

func TestNewestInFamily(t *testing.T) {
	candidates := stableCandidates([]string{
		"3.9.5", "4.2.0", "4.10.4", "4.9.9", "4.11.0-beta.1", "5.0.0", "0.2.28", "0.2.29", "0.3.0", "4",
	})
	cases := map[string]string{"4": "4.10.4", "3": "3.9.5", "0.2": "0.2.29", "0": "0.3.0"}
	for family, want := range cases {
		parsed, _ := parseNumericVersion(family)
		got, ok := newestInFamily(parsed, candidates)
		if !ok || got != want {
			t.Fatalf("newestInFamily(%s) = %q, %v; want %q", family, got, ok, want)
		}
	}
	if got, ok := newestInFamily([]int{6}, candidates); ok {
		t.Fatalf("newestInFamily(6) = %q, want no match", got)
	}
}

func withRegistryServer(t *testing.T, routes map[string]string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := routes[r.URL.EscapedPath()]
		if !ok {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv
}

func resolveForTest(t *testing.T, raw string) string {
	t.Helper()
	spec, err := parseSDKVersionSpec(raw)
	if err != nil {
		t.Fatal(err)
	}
	sdkResolutions.Delete(spec.raw)
	t.Cleanup(func() { sdkResolutions.Delete(spec.raw) })
	got, err := resolveSDKVersion(spec, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("resolveSDKVersion(%s): %v", raw, err)
	}
	return got
}

func TestResolveSDKVersionRegistries(t *testing.T) {
	srv := withRegistryServer(t, map[string]string{
		"/@launchdarkly%2Fjs-client-sdk":                                              `{"versions":{"4.10.3":{},"4.10.4":{},"5.0.0-beta.1":{},"3.0.0":{}}}`,
		"/api/packages/launchdarkly_flutter_client_sdk":                               `{"versions":[{"version":"4.20.3"},{"version":"4.21.0"},{"version":"4.22.0","retracted":true}]}`,
		"/package/launchdarkly-server-sdk/preferred":                                  `{"normal-version":["4.6.0","4.5.1","3.1.1"],"deprecated-version":["4.7.0"]}`,
		"/api/packages/launchdarkly_server_sdk":                                       `{"releases":[{"version":"3.11.2"},{"version":"3.12.0"}],"retirements":{"3.12.0":{"reason":"invalid"}}}`,
		"/manifests/launchdarkly/manifest-5.3.json":                                   `{"repository":{"launchdarkly-server-sdk":{"2.1.3-0":[],"2.2.0-0":[],"2.2.0-1":[]}}}`,
		"/maven2/com/launchdarkly/launchdarkly-android-client-sdk/maven-metadata.xml": `<metadata><versioning><versions><version>5.15.0</version><version>5.16.0</version><version>6.0.0-rc1</version></versions></versioning></metadata>`,
	})
	for _, v := range []*string{&npmRegistryURL, &pubRegistryURL, &hackageURL, &hexURL, &luarocksURL} {
		old := *v
		*v = srv.URL
		t.Cleanup(func() { *v = old })
	}
	oldMaven := mavenRepositoryURLs
	mavenRepositoryURLs = []string{srv.URL + "/missing", srv.URL + "/maven2"}
	t.Cleanup(func() { mavenRepositoryURLs = oldMaven })

	cases := map[string]string{
		"npm:@launchdarkly/js-client-sdk@4":                        "4.10.4",
		"pub:launchdarkly_flutter_client_sdk@4":                    "4.21.0",
		"hackage:launchdarkly-server-sdk@4":                        "4.6.0",
		"hex:launchdarkly_server_sdk@3":                            "3.11.2",
		"luarocks:launchdarkly/launchdarkly-server-sdk@2":          "2.2.0-1",
		"maven:com.launchdarkly:launchdarkly-android-client-sdk@5": "5.16.0",
	}
	for raw, want := range cases {
		if got := resolveForTest(t, raw); got != want {
			t.Fatalf("resolve %s = %q, want %q", raw, got, want)
		}
	}
}

func TestResolveSDKVersionGitHubTags(t *testing.T) {
	old := gitLsRemote
	gitLsRemote = func(repoURL, pattern string) ([]byte, error) {
		if repoURL != "https://github.com/launchdarkly/cpp-sdks" || pattern != "refs/tags/launchdarkly-cpp-server-v*" {
			t.Fatalf("gitLsRemote(%q, %q)", repoURL, pattern)
		}
		return []byte("a\trefs/tags/launchdarkly-cpp-server-v3.9.1\n" +
			"b\trefs/tags/launchdarkly-cpp-server-v3.14.0\n" +
			"c\trefs/tags/launchdarkly-cpp-server-v4.0.0-rc.1\n" +
			"d\trefs/tags/launchdarkly-cpp-server-redis-source-v2.2.0\n"), nil
	}
	t.Cleanup(func() { gitLsRemote = old })

	if got := resolveForTest(t, "github-tag:launchdarkly/cpp-sdks:launchdarkly-cpp-server-v@3"); got != "launchdarkly-cpp-server-v3.14.0" {
		t.Fatalf("resolved tag = %q", got)
	}
}

func TestBuildArgsResolvesSDKFamiliesAndChangesImageTag(t *testing.T) {
	versions := `{"versions":{"4.10.3":{}}}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(versions))
	}))
	t.Cleanup(srv.Close)
	old := npmRegistryURL
	npmRegistryURL = srv.URL
	t.Cleanup(func() { npmRegistryURL = old })

	root := t.TempDir()
	writeValidatorVersionFiles(t, root, "IMAGE=node:22\n", "NPM_VERSION=12\n", "GO_VERSION=1.25.12\n",
		"LD_SDK_VERSION=npm:@launchdarkly/js-client-sdk@4\n")
	runnerDir := filepath.Join(root, "languages", "test")
	if err := os.MkdirAll(runnerDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(runnerDir, "Dockerfile"), []byte("ARG IMAGE\nFROM ${IMAGE}\nARG LD_SDK_VERSION\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	resolve := func() ([]string, string) {
		sdkResolutions.Delete("npm:@launchdarkly/js-client-sdk@4")
		var out bytes.Buffer
		args, err := BuildArgs(root, "test", &out)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(out.String(), "resolved npm:@launchdarkly/js-client-sdk@4 -> ") {
			t.Fatalf("resolution not logged: %q", out.String())
		}
		tag, err := validatorImageTag(root, runnerDir, "test", args)
		if err != nil {
			t.Fatal(err)
		}
		return args, tag
	}
	t.Cleanup(func() { sdkResolutions.Delete("npm:@launchdarkly/js-client-sdk@4") })

	args, firstTag := resolve()
	if want := []string{"IMAGE=node:22", "LD_SDK_VERSION=4.10.3"}; !reflect.DeepEqual(args, want) {
		t.Fatalf("BuildArgs() = %#v, want %#v", args, want)
	}
	versions = `{"versions":{"4.10.3":{},"4.10.4":{}}}`
	args, secondTag := resolve()
	if args[1] != "LD_SDK_VERSION=4.10.4" {
		t.Fatalf("BuildArgs() after release = %#v", args)
	}
	if firstTag == secondTag {
		t.Fatalf("image tag %s did not change after a new release", firstTag)
	}
}
