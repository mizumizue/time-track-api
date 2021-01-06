package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type HttpClient interface {
	Do(req *http.Request) ([]byte, error)
}

type HttpCustomClient struct {
	client *http.Client
}

func NewHttpCustomClient(client *http.Client) HttpClient {
	return &HttpCustomClient{client: client}
}

func (customClient *HttpCustomClient) Do(req *http.Request) ([]byte, error) {
	res, err := customClient.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request error. http code: %d, status: %s", res.StatusCode, res.Status)
	}
	return body, nil
}

func NewRequestWithHeader(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TOKEN")))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
