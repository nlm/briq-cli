package briq

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	hc      *http.Client
	headers map[string][]string
}

func NewClient(token string) (*Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Jar:     cookieJar,
		Timeout: 10 * time.Second,
	}
	return &Client{
		hc: &client,
		headers: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {fmt.Sprintf("Bearer %s", token)},
			"User-Agent":    {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36"},
		},
	}, nil
}

func (client *Client) do(ctx context.Context, method, targetUrl string, request, response any) error {
	url, err := url.Parse(targetUrl)
	if err != nil {
		return err
	}
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req := &http.Request{
		Method: method,
		Header: client.headers,
		URL:    url,
		Body:   io.NopCloser(bytes.NewReader(data)),
	}
	res, err := client.hc.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New(res.Status)
	}
	rdata, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(rdata, response)
}
