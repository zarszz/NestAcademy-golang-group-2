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
	data, err := c.build(http.MethodGet, path, nil)
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

	data, err := c.build(http.MethodPost, path, dataByte)
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

func (c *Client) build(method string, path string, payload []byte) (interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseUrl, path), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

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
