package requests

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"net/http"
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

func readerFromBody(body *map[string]string) *bytes.Buffer {
	value, _ := json.Marshal(body)
	return bytes.NewBuffer(value)
}

func setReqHeaders(req *http.Request, headers *map[string]string) {
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
}