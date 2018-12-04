package requests

import (
	"bytes"
	"fmt"
	"net/http"
)

func formatURL(host string, path string, query *map[string]string) string {
	url := host + path
	if query != nil {
		url = url + "?"
		for key, value := range *query {
			url = fmt.Sprintf("%s%s=%s&", url, key, value)
		}
	}
	return url
}

func bufferFromBody(body *map[string]string)(*bytes.Buffer, error)  {
	buffer := new(bytes.Buffer)
	for key, value := range *body {
		_, err := fmt.Fprintf(buffer, "%s=\"%s\"\n", key, value)
		if err != nil {
			return nil, err
		}
	}
	return buffer, nil
}

func setReqHeaders(req *http.Request, headers *map[string]string) {
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
}