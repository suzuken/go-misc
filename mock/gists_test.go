package mock

import (
	"io"
	"strings"
	"testing"
)

// パターン1
func TestGetGists(t *testing.T) {
	doGistsRequest = func(user string) (io.Reader, error) {
		return strings.NewReader(`
[
	{"html_url": "https://gist.github.com/example1"},
	{"html_url": "https://gist.github.com/example2"}
]
		`), nil
	}
	urls, err := ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}

type dummyDoer struct{}

func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
[
	{"html_url": "https://gist.github.com/example1"},
	{"html_url": "https://gist.github.com/example2"}
]
	`), nil
}

// パターン1
func TestGetGists2(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
	for _, u := range urls {
		t.Logf("%v", u)
	}
}
