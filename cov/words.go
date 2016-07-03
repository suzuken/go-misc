package cov

import (
	"strings"
)

func Words(s string) string {
	c := len(strings.Fields(s))
	switch {
	case c == 0:
		return "wordless?"
	case c == 1:
		return "one word"
	case c < 4:
		return "a few words"
	case c < 8:
		return "many words"
	default:
		return "too many words"
	}
}
