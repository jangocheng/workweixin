package cores

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	err     error
	client  *http.Client
	request *http.Request
}

func InitClient(method string, url string, ioBody io.Reader) *HttpClient {
	c := &HttpClient{
		client: &http.Client{},
	}
	req, err := http.NewRequest(method, url, ioBody)
	if err != nil {
		c.err = err
		return c
	}
	c.request = req
	return c
}

func (c *HttpClient) HttpDo(v interface{}) error {
	if c.err != nil {
		return c.err
	}
	resp, err := c.client.Do(c.request)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}
