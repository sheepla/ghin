//nolint
package gh

import (
	"testing"

	loader "github.com/peteole/testdata-loader"
)

func TestReleasesURL(t *testing.T) {
	want := "https://api.github.com/repos/sheepla/srss/releases"
	have := releasesURL("sheepla", "srss").String()
	if have != want {
		t.Fatalf("have=%s, want=%s", have, want)
	}
}

func TestParseAsReleasesData(t *testing.T) {
	bytes := loader.GetTestFile("_testdata/releases.json")
	releases, err := parseAsReleasesData(bytes)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range *releases {
		t.Logf(
			"[%s]\n- Created at %s\n- Published at %s\n- Page URL: %s\n-----------\n%s\n",
			item.Tag,
			item.CreatedAt,
			item.PublishedAt,
			item.PageURL,
			item.ReleaseNote,
		)
	}
}

func TestParseAsAssetsData(t *testing.T) {
	bytes := loader.GetTestFile("_testdata/releases.json")
	releases, err := parseAsReleasesData(bytes)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range *releases {
		for _, a := range *item.Assets {
			t.Logf(
				"%s (%d downloads)\nSize:%d, ContentType:%s\nDownload from: %s\n\n",
				a.Name,
				a.DownloadCount,
				a.Size,
				a.ContentType,
				a.DownloadURL,
			)
		}
	}
}
