package osugosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *HTTPClient) Request(method, path string, query url.Values, body, result any) error {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return err
	}

	u.RawQuery = query.Encode()

	var reader io.Reader
	if body != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
		reader = buf
	}

	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed: %s", resp.Status)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

func (c *HTTPClient) Get(path string, query url.Values, result any) error {
	return c.Request(http.MethodGet, path, query, nil, result)
}

func (c *HTTPClient) Post(path string, query url.Values, body, result any) error {
	return c.Request(http.MethodPost, path, query, body, result)
}

func (c *HTTPClient) Put(path string, query url.Values, body, result any) error {
	return c.Request(http.MethodPut, path, query, body, result)
}

func (c *HTTPClient) Patch(path string, query url.Values, body, result any) error {
	return c.Request(http.MethodPatch, path, query, body, result)
}

func (c *HTTPClient) Delete(path string, query url.Values, body, result any) error {
	return c.Request(http.MethodDelete, path, query, body, result)
}
