package ui

import (
	"fmt"

	"github.com/dustin/go-humanize"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/ghin/gh"
)

func SelectRepo(result []gh.SearchResult) (int, error) {
	//nolint:wrapcheck
	return fzf.Find(
		result,
		func(i int) string {
			return fmt.Sprintf("%s/%s", result[i].Owner, result[i].Name)
		},
		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return "NO RESULTS"
			}

			return fmt.Sprintf(
				"%s/%s (%d stars, %d forks)\n\n%s\n\nlicense: %s\nlanguage: %s\ncreated at: %s\nupdated at: %s",
				result[i].Owner,
				result[i].Name,
				result[i].Stars,
				result[i].Forks,
				result[i].Description,
				result[i].License,
				result[i].Language,
				humanize.Time(result[i].CreatedAt),
				humanize.Time(result[i].UpdatedAt),
			)
		}),
	)
}

func SelectTag(releases []gh.Release) (int, error) {
	//nolint:wrapcheck
	return fzf.Find(
		releases,
		func(i int) string {
			return releases[i].Tag
		},

		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return "NO RESULTS"
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

func SelectAsset(assets []gh.Asset) (int, error) {
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
