package validate

import (
	"testing"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
)

func TestCanonicalEnv(t *testing.T) {
	// Order-independent: the same map always encodes identically.
	a := canonicalEnv(map[string]string{"B": "2", "A": "1"})
	b := canonicalEnv(map[string]string{"A": "1", "B": "2"})
	if a != b {
		t.Fatalf("canonicalEnv not order-independent: %q vs %q", a, b)
	}
	// Distinct values produce distinct keys.
	if canonicalEnv(map[string]string{"A": "1"}) == canonicalEnv(map[string]string{"A": "2"}) {
		t.Fatal("canonicalEnv collapsed distinct values")
	}
	// A key=value pair can't be confused with the absence of a key
	// (NUL separators prevent "A"+"=BC" colliding with "AB"+"=C").
	if canonicalEnv(map[string]string{"A": "BC"}) == canonicalEnv(map[string]string{"AB": "C"}) {
		t.Fatal("canonicalEnv ambiguous encoding")
	}
}

func TestGroupUnits(t *testing.T) {
	unit := func(id, runtime string, env map[string]string) *resolvedUnit {
		return &resolvedUnit{
			snippet:  &model.Snippet{Frontmatter: model.Frontmatter{ID: id}},
			runtime:  runtime,
			extraEnv: env,
		}
	}
	def := map[string]string{"SNIPPET_CHECK": "runtime"}
	v1 := map[string]string{"SNIPPET_CHECK": "runtime", "LD_RUST_SDK_VERSION": "1"}

	units := []*resolvedUnit{
		unit("a", "rust", def),
		unit("b", "rust", v1),
		unit("c", "rust", def),
		unit("d", "go", def),
	}
	groups, order := groupUnits(units)

	// Three buckets: rust+default, rust+v1, go+default.
	if len(groups) != 3 {
		t.Fatalf("want 3 groups, got %d (%v)", len(groups), order)
	}
	// Same (runtime, env) snippets land together.
	var rustDefault []*resolvedUnit
	for key, g := range groups {
		if len(g) == 2 {
			rustDefault = g
			_ = key
		}
	}
	if rustDefault == nil {
		t.Fatal("expected a 2-snippet rust/default bucket")
	}
	// Order is deterministic (sorted) so logs are stable.
	for i := 1; i < len(order); i++ {
		if order[i-1] > order[i] {
			t.Fatalf("group order not sorted: %v", order)
		}
	}
}
