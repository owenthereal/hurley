package hurley

import (
	"net/http"
)

type Handler interface {
	PrepareRequest(req *http.Request) error
	PrepareResponse(resp *http.Response) error
}

type Client struct {
	client   *http.Client
	Handlers []Handler
}

func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	for _, h := range c.Handlers {
		err = h.PrepareRequest(req)
		if err != nil {
			return
		}
	}

	resp, err = c.client.Do(req)
	if err != nil {
		return
	}

	for _, h := range c.Handlers {
		err = h.PrepareResponse(resp)
		if err != nil {
			return
		}
	}

	return
}

func (c *Client) Use(handler Handler) {
	c.Handlers = append(c.Handlers, handler)
}

func New() *Client {
	return &Client{client: &http.Client{}, Handlers: make([]Handler, 0)}
}
