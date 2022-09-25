//nolint
package ui

import (
	"testing"
	"time"

	"github.com/sheepla/ghin/gh"
)

/*
func TestFindRepo(t *testing.T) {
	result := []gh.SearchResult{
		{
			Owner:       "sheepla",
			Name:        "myrepo",
			License:     "MIT",
			Description: "my super cool repository",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Size:        1234,
			Language:    "Go",
			Stars:       123456,
			Forks:       1234567,
		},
		{
			Owner:       "sheepla",
			Name:        "mytool",
			License:     "MIT",
			Description: "my ultimate cool tool",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Size:        4321,
			Language:    "Go",
			Stars:       54321,
			Forks:       654321,
		},
	}

	idx, err := SelectRepo(result)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(idx)
}
*/

func TestSelectTag(t *testing.T) {
	releases := []gh.Release{
		{
			Tag:         "v0.0.1",
			PageURL:     "https://.....",
			CreatedAt:   time.Now(),
			PublishedAt: time.Now(),
			ReleaseNote: "What's new!",
			Author:      "sheepla",
			Assets: []gh.Asset{
				{
					Name:          "foo.gz",
					Size:          12345,
					Uploader:      "bot",
					ContentType:   "application/gzip",
					DownloadCount: 10,
					DownloadURL:   "https://.....",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				},
				{
					Name:          "bar.zip",
					Size:          12345,
					Uploader:      "bot",
					ContentType:   "application/zip",
					DownloadCount: 10,
					DownloadURL:   "https://.....",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				},
			},
		},
		{
			Tag:         "v0.0.2",
			PageURL:     "https://.....",
			CreatedAt:   time.Now(),
			PublishedAt: time.Now(),
			ReleaseNote: "What's new!",
			Author:      "sheepla",
			Assets: []gh.Asset{
				{
					Name:          "foo.gz",
					Size:          12345,
					Uploader:      "bot",
					ContentType:   "application/gzip",
					DownloadCount: 10,
					DownloadURL:   "https://.....",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				},
				{
					Name:          "bar.zip",
					Size:          12345,
					Uploader:      "bot",
					ContentType:   "application/zip",
					DownloadCount: 10,
					DownloadURL:   "https://.....",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				},
			},
		},
	}

	idx, err := SelectTag(releases)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(idx)
}
