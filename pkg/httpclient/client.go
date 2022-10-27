package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	data, err := c.build(http.MethodGet, path, nil, nil, nil)
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
	data, err := c.build(http.MethodGet, path, nil, headers, query)
	if err != nil {
		return nil, err
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataByte, nil
}

func (c *Client) Post(path string, payload interface{}) ([]byte, error) {
	dataByte, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data, err := c.build(http.MethodPost, path, dataByte, nil, nil)
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

func (c *Client) build(method string, path string, payload []byte, headers *map[string]string, query *map[string]string) (interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseUrl, path), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	q := req.URL.Query()
	if query != nil {
		for k, v := range *query {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		return nil, errors.New("not found")
	}

	var data interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
