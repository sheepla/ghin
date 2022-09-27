package gh

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/tidwall/gjson"
)

const timeout = 10

type SearchParam struct {
	Query string
	// Language     string
	// User         string
	// Organization string
	// Followers    int
	// CreatedAt    *time.Duration
	// UpdatedAt    *time.Duration
	// Topic        string
	// TopicsCount  string
	// License      string
	// Mirror       bool
	// HasArchved   bool
}

type SearchResult struct {
	Owner       string
	Name        string
	License     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Size        int64
	Language    string
	Stars       int64
	Forks       int64
}

func NewSearchParam(query string) *SearchParam {
	return &SearchParam{
		Query: query,
	}
}

func (param *SearchParam) toURL() *url.URL {
	//nolint:exhaustivestruct,exhaustruct,varnamelen
	u := &url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   "search/repositories",
	}
	q := u.Query()
	q.Add("q", param.toQueryString())

	u.RawQuery = q.Encode()

	return u
}

func (param *SearchParam) toQueryString() string {
	return fmt.Sprintf("%s fork:false", param.Query)
}

func Search(param *SearchParam) (*[]SearchResult, error) {
	body, err := fetch(param.toURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the result: %w", err)
	}
	defer body.Close()

	content, err := io.ReadAll(body)
	if err != nil {
		//nolint:goerr113
		return nil, errors.New("failed to read response body")
	}

	results, err := parseAsSearchResults(content)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return results, nil
}

func parseAsSearchResults(bytes []byte) (*[]SearchResult, error) {
	if !gjson.ValidBytes(bytes) {
		//nolint:goerr113
		return nil, errors.New("invalid JSON format")
	}

	//nolint:prealloc
	var results []SearchResult

	items := gjson.GetBytes(bytes, "items")

	for _, item := range items.Array() {
		result := SearchResult{
			Owner:       item.Get("owner.login").String(),
			Name:        item.Get("name").String(),
			Description: item.Get("description").String(),
			License:     item.Get("license").Get("name").String(),
			Size:        item.Get("size").Int(),
			Stars:       item.Get("stargazers_count").Int(),
			Forks:       item.Get("forks_count").Int(),
			Language:    item.Get("language").String(),
			CreatedAt:   item.Get("created_at").Time(),
			UpdatedAt:   item.Get("updated_at").Time(),
		}

		results = append(results, result)
	}

	return &results, nil
}
