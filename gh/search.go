package gh

import (
	"errors"
	"fmt"
	"io"
	"net/http"
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
		return nil, errors.New("failed to read response body")
	}

	results, err := parseAsSearchResults(content)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return results, nil
}

//nolint:interfacer
func fetch(url *url.URL) (io.ReadCloser, error) {
	//nolint:noctx
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %w", err)
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	//nolint:exhaustivestruct,exhaustruct
	c := &http.Client{
		Timeout: timeout * time.Second,
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the result: %w", err)
	}

	if res.StatusCode < 200 || 300 <= res.StatusCode {
		//nolint:goerr113
		return nil, fmt.Errorf("http status error: %s", res.Status)
	}

	return res.Body, nil
}

func parseAsSearchResults(content []byte) (*[]SearchResult, error) {
	if !gjson.ValidBytes(content) {
		//nolint:goerr113
		return nil, errors.New("invalid JSON format")
	}

	//nolint:prealloc
	var results []SearchResult

	items := gjson.GetBytes(content, "items")

	for _, item := range items.Array() {
		result := SearchResult{
			Owner:       gjson.Get(items.String(), "owner.login").String(),
			Name:        gjson.Get(item.String(), "name").String(),
			Description: gjson.Get(item.String(), "description").String(),
			License:     gjson.Get(item.String(), "license").String(),
			Size:        gjson.Get(item.String(), "size").Int(),
			Stars:       gjson.Get(item.String(), "stargazers_count").Int(),
			Forks:       gjson.Get(item.String(), "forks_count").Int(),
			Language:    gjson.Get(item.String(), "language").String(),
			CreatedAt:   gjson.Get(item.String(), "created_at").Time(),
			UpdatedAt:   gjson.Get(item.String(), "updated_at").Time(),
		}

		results = append(results, result)
	}

	return &results, nil
}
