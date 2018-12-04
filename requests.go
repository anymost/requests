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
	buffer, err := bufferFromBody(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("GET", url, buffer)
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
	buffer, err := bufferFromBody(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PATCH", url, buffer)
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
	buffer, err := bufferFromBody(body)
	request, err := http.NewRequest("DELETE", url, buffer)
	if err != nil {
		return nil, err
	}
	setReqHeaders(request, headers)
	resp, err := client.Do(request)
	return resp, err
}