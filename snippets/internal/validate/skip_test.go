package validate

import "testing"

func TestSkipID(t *testing.T) {
	cases := []struct {
		list, id string
		want     bool
	}{
		{"", "a/b/c", false},
		{"a/b/c", "a/b/c", true},
		{"a/b/c", "a/b/d", false},
		{"a/b/c,x/y/z", "x/y/z", true},
		{"a/b/c, x/y/z", "x/y/z", true},
		{"a/b/c,x/y/z", "a/b/c", true},
		{"a/b/c,x/y/z", "q/r/s", false},
	}
	for _, c := range cases {
		if got := skipID(c.list, c.id); got != c.want {
			t.Errorf("skipID(%q, %q) = %v, want %v", c.list, c.id, got, c.want)
		}
	}
}
