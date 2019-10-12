package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, s := range cases {
		got := Reverse(s.in)
		if got != s.want {
			t.Errorf("Reverse(%q) == %q, want %q", s.in, got, s.want)
		}
	}
}
