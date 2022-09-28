package gh

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"time"

	"github.com/tidwall/gjson"
)

type Release struct {
	Tag         string
	PageURL     string
	CreatedAt   time.Time
	PublishedAt time.Time
	ReleaseNote string
	Author      string
	Assets      []Asset
}

type Releases []Release

type Asset struct {
	Name          string
	Size          int64
	Uploader      string
	ContentType   string
	DownloadCount int64
	DownloadURL   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Assets []Asset

func GetReleases(owner, repo string) (*Releases, error) {
	url := releasesURL(owner, repo)

	body, err := fetch(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the result: %w", err)
	}

	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		//nolint:goerr113
		return nil, errors.New("failed to read response body")
	}

	releases, err := parseAsReleasesData(content)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return releases, nil
}

func releasesURL(owner, repo string) *url.URL {
	//nolint:exhaustivestruct,exhaustruct
	return &url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   path.Join("repos", owner, repo, "releases"),
	}
}

func parseAsReleasesData(bytes []byte) (*Releases, error) {
	if !gjson.ValidBytes(bytes) {
		//nolint:goerr113
		return nil, errors.New("invalid JSON format")
	}

	data := gjson.ParseBytes(bytes)

	//nolint:prealloc
	var releases Releases

	for _, item := range data.Array() {
		rel := Release{
			Tag:         item.Get("name").String(),
			PageURL:     item.Get("html_url").String(),
			CreatedAt:   item.Get("created_at").Time(),
			PublishedAt: item.Get("published_at").Time(),
			Author:      item.Get("author").Get("login").String(),
			ReleaseNote: item.Get("body").String(),
			Assets:      parseAsAssetsData(item),
		}

		releases = append(releases, rel)
	}

	return &releases, nil
}

func parseAsAssetsData(data gjson.Result) Assets {
	//nolint:prealloc
	var assets []Asset

	//nolint:exhaustivestruct,exhaustruct
	for _, item := range data.Get("assets").Array() {
		ast := Asset{
			Name:          item.Get("name").String(),
			Uploader:      item.Get("uploader").Get("login").String(),
			ContentType:   item.Get("content_type").String(),
			Size:          item.Get("size").Int(),
			DownloadCount: item.Get("download_count").Int(),
			DownloadURL:   item.Get("browser_download_url").String(),
		}

		assets = append(assets, ast)
	}

	return assets
}
