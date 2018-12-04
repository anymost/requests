package requests

import (
	"net/http"
)

type Requests struct {
	Host string
}


func (req *Requests) GET(path string, headers *map[string]string,  query *map[string]string) (*http.Response, error)  {
	client := &http.Client{}
	url := formatURL(req.Host, path, query)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setReqHeaders(request, headers)
	resp, err := client.Do(request)
	return resp, err
}

func (req *Requests) POST(path string, headers *map[string]string, query *map[string]string, body *map[string]string)(*http.Response, error)  {
	client := &http.Client{}
	url := formatURL(req.Host, path, query)
	reader := readerFromBody(body)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	setReqHeaders(request, headers)
	resp, err := client.Do(request)
	return resp, err
}

func (req *Requests) PATCH(path string, headers *map[string]string, query *map[string]string, body *map[string]string)(*http.Response, error)  {
	client := &http.Client{}
	url := formatURL(req.Host, path, query)
	reader:= readerFromBody(body)
	request, err := http.NewRequest("PATCH", url, reader)
	if err != nil {
		return nil, err
	}
	setReqHeaders(request, headers)
	resp, err := client.Do(request)
	return resp, err
}

func (req *Requests) DELETE(path string, headers *map[string]string, query *map[string]string, body *map[string]string)(*http.Response, error)  {
	client := &http.Client{}
	url := formatURL(req.Host, path, query)
	reader := readerFromBody(body)
	request, err := http.NewRequest("DELETE", url, reader)
	if err != nil {
		return nil, err
	}
	setReqHeaders(request, headers)
	resp, err := client.Do(request)
	return resp, err
}