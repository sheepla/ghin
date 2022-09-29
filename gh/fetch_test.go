//nolint:testpackage,paralleltest
package gh

import (
	"io"
	"net/url"
	"testing"
)

func TestFetch(t *testing.T) {
	u, err := url.Parse("https://api.github.com/search/repositories?q=vim+fork%3Afalse")
	if err != nil {
		t.Fatal(u)
	}

	body, err := fetch(u)
	if err != nil {
		t.Fatal(err)
	}

	defer body.Close()

	bytes, err := io.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}
