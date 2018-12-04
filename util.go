package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func formatURL(host string, path string, query *map[string]string) string {
	fullURL := host + path
	if query != nil {
		fullURL = fullURL + "?"
		for key, value := range *query {
			fullURL = fmt.Sprintf("%s%s=%s&", fullURL, key, value)
		}
	}
	return fullURL
}

func readerFromBody(body *map[string]string) *strings.Reader {
	data := url.Values{}
	for key, value := range *body {
		data.Set(key, value)
	}
	reader := strings.NewReader(data.Encode())
	return reader
}

func setReqHeaders(req *http.Request, headers *map[string]string) {
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
}