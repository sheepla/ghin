//nolint
package gh

import (
	"testing"

	loader "github.com/peteole/testdata-loader"
)

var param = &SearchParam{
	Query: "vim",
}

func TestSearchParamToQuery(t *testing.T) {
	if q := param.toQueryString(); q != "vim fork:false" {
		t.Log(q)
		t.Fatal(q)
	}
}

func TestSearchParamToURL(t *testing.T) {
	if u := param.toURL().String(); u != "https://api.github.com/search/repositories?q=vim+fork%3Afalse" {
		t.Log(u)
		t.Fatalf("wrong URL: %s", u)
	}
}

func TestParseAsSearchResults(t *testing.T) {
	bytes := loader.GetTestFile("_testdata/search.json")

	results, err := parseAsReposData(bytes)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range *results {
		t.Logf(
			"%s/%s %dstars, %dforks\n%s\n\n",
			item.Owner,
			item.Name,
			item.Stars,
			item.Forks,
			item.Description,
		)
	}
}
