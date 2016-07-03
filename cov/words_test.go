package cov

import (
	"testing"
)

type Case struct {
	in, out string
}

var cases = []Case{
	{"今日は天気ですね", "one word"},
	{"Go Go Go Go Go Go Go!", "many words"},
}

func TestWords(t *testing.T) {
	for i, c := range cases {
		w := Words(c.in)
		if w != c.out {
			t.Errorf("#%d: Words(%s) got %s; want %s", i, c.in, w, c.out)
		}
	}
}
