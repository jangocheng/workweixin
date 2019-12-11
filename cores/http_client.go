package cores

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Err      error
	Client   *http.Client
	Request  *http.Request
}

func InitClient(method string, url string, ioBody io.Reader) *HttpClient {
	c := &HttpClient{
		Client: &http.Client{},
	}
	req, err := http.NewRequest(method, url, ioBody)
	if err != nil {
		c.Err = err
		return c
	}
	c.Request = req
	return c
}

func (c *HttpClient) HttpDo() ([]byte, error) {
	if c.Err != nil {
		return nil, c.Err
	}
	resp, err := c.Client.Do(c.Request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *HttpClient) HttpResult(v interface{}) error {
	body, err := c.HttpDo()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}
