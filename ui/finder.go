package ui

import (
	"fmt"

	"github.com/dustin/go-humanize"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/ghin/gh"
)

const noResultsText = "NO RESULTS"

//nolint:wrapcheck
func SelectRepo(repos gh.Repos) (*gh.Repo, error) {
	idx, err := fzf.Find(
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

	return &repos[idx], err
}

//nolint:wrapcheck
func SelectTag(rels gh.Releases) (*gh.Release, error) {
	idx, err := fzf.Find(
		rels,
		func(i int) string {
			return rels[i].Tag
		},

		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return noResultsText
			}

			return fmt.Sprintf(
				"%s\n\nby %s\n\ncreated at: %s, published at: %s\n\nSee release page: %s\n\n─────────\n\n%s\n\n─────────\n\n",
				rels[i].Tag,
				rels[i].Author,
				humanize.Time(rels[i].CreatedAt),
				humanize.Time(rels[i].PublishedAt),
				rels[i].PageURL,
				rels[i].ReleaseNote,
			)
		}),
	)

	return &rels[idx], err
}

//nolint:wrapcheck
func SelectAsset(assets gh.Assets) (*gh.Asset, error) {
	idx, err := fzf.Find(
		assets,
		func(i int) string {
			size := humanize.Bytes(uint64(assets[i].Size))

			return fmt.Sprintf("%s (%s)", assets[i].Name, size)
		},

		//nolint:varnamelen
		fzf.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return noResultsText
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

	return &assets[idx], err
}
