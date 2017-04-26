package stringcleansing

import "testing"

func TestReplaceHalfHyphen(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"090ー1234-5678", "090-1234-5678"},
		{"080ー1234—5678", "080-1234-5678"},
		{"070-1234-5678", "070-1234-5678"},
		{"", ""},
	}

	for _, c := range cases {
		if got := ReplaceHalfHyphen(c.in); got != c.want {
			t.Errorf("ReplaceHalfHyphen(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
