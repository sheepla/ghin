package ui

import (
	"fmt"

	"github.com/dustin/go-humanize"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/ghin/gh"
)

const noResultsText = "NO RESULTS"

func SelectRepo(repos gh.Repos) (int, error) {
	//nolint:wrapcheck
	return fzf.Find(
		repos,
		func(i int) string {
			return fmt.Sprintf("%s/%s", repos[i].Owner, repos[i].Name)
		},
		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return noResultsText
			}

			return fmt.Sprintf(
				"%s/%s (%d stars, %d forks)\n\n%s\n\nlicense: %s\nlanguage: %s\ncreated at: %s\nupdated at: %s",
				repos[i].Owner,
				repos[i].Name,
				repos[i].Stars,
				repos[i].Forks,
				repos[i].Description,
				repos[i].License,
				repos[i].Language,
				humanize.Time(repos[i].CreatedAt),
				humanize.Time(repos[i].UpdatedAt),
			)
		}),
	)
}

func SelectTag(releases gh.Releases) (int, error) {
	//nolint:wrapcheck
	return fzf.Find(
		releases,
		func(i int) string {
			return releases[i].Tag
		},

		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return noResultsText
			}

			return fmt.Sprintf(
				"%s\n\nby %s\n\ncreated at: %s, published at: %s\n\nSee release page: %s\n\n─────────\n\n%s\n\n─────────\n\n",
				releases[i].Tag,
				releases[i].Author,
				humanize.Time(releases[i].CreatedAt),
				humanize.Time(releases[i].PublishedAt),
				releases[i].PageURL,
				releases[i].ReleaseNote,
			)
		}),
	)
}

func SelectAsset(assets gh.Assets) (int, error) {
	//nolint:wrapcheck
	return fzf.Find(
		assets,
		func(i int) string {
			size := humanize.Bytes(uint64(assets[i].Size))

			return fmt.Sprintf("%s (%s)", assets[i].Name, size)
		},

		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return "NO RESULTS"
			}

			size := humanize.Bytes(uint64(assets[i].Size))

			return fmt.Sprintf(
				"%s (%s)\n\ncontent type: %s\n\nuploaded by %s\n",
				assets[i].Name,
				size,
				assets[i].ContentType,
				assets[i].Uploader,
			)
		}),
	)
}
