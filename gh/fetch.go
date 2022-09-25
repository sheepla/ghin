package gh

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

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
