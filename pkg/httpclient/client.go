package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	baseUrl string
}

func NewHttpClient(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
	}
}

func (c *Client) Get(path string) ([]byte, error) {
	data, err := c.build(http.MethodGet, path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataByte, nil
}

func (c *Client) GetWithHeadersAndQuery(path string, headers *map[string]string, query *map[string]string) ([]byte, error) {
	data, err := c.build(http.MethodGet, path, nil, headers, query, nil)
	if err != nil {
		return nil, err
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataByte, nil
}

func (c *Client) PostWithHeader(path string, headers *map[string]string, payload *string) ([]byte, error) {
	contentType := "application/x-www-form-urlencoded"
	data, err := c.build(http.MethodPost, path, payload, headers, nil, &contentType)
	if err != nil {
		return nil, err
	}
	respData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func (c *Client) Post(path string, payload *string) ([]byte, error) {

	data, err := c.build(http.MethodPost, path, payload, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	respData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return respData, nil
}

func (c *Client) build(method string, path string, payload *string, headers *map[string]string, query *map[string]string, contentType *string) (interface{}, error) {
	client := http.Client{}

	var req *http.Request

	if payload != nil {
		req, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseUrl, path), strings.NewReader(*payload))
	} else {
		req, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseUrl, path), nil)
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	if query != nil {
		q := req.URL.Query()
		for k, v := range *query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if contentType == nil {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", *contentType)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		return nil, errors.New("not found")
	}

	fmt.Println(resp.Body)

	var data interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
