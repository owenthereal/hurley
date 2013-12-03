Hurley
======

Hurley and Faraday are both from the [lost island](lostisland).

```go
package main

import "github.com/jingweno/hurley"

type TestHandler struct {
}

func (h *TestHandler) PrepareRequest(req *http.Request) error {
	req.Header.Add("Content-Type", "application/json")
	return nil
}

func (h *TestHandler) PrepareResponse(resp *http.Response) error {
	resp.Status = "201"
	return nil
}

func main() {
  c := hurley.New()
  c.Use(&TestHandler{})
  resp, err := c.Get("https://api.github.com/repos/jingweno/hurley")
}
```

See [test](https://github.com/jingweno/hurley/blob/master/hurley_test.go) as an example.
