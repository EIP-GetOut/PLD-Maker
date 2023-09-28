package tools

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

// Description: Request with GET Method
//
// Parameters:
//   - *http.Client
//   - Url(string)
//   - data([]byte)
//   - header(...tools.Pair) (params-pack you can give how much header you want)
func RequestGet(client *http.Client, url string, header url.Values) ([]byte, error) {
	return Request(client, http.MethodGet, url, nil, header)
}

// Description: Request with POST Method
//
// Parameters:
//   - *http.Client
//   - Url(string)
//   - data([]byte)
//   - header(...tools.Pair) (params-pack you can give how much header you want)
func RequestPost(client *http.Client, url string, data []byte, header url.Values) ([]byte, error) {
	return Request(client, http.MethodPost, url, data, header)
}

// Description: Request with DELETE Method
//
// Parameters:
//   - *http.Client
//   - Url(string)
//   - data([]byte)
//   - header(...tools.Pair) (params-pack you can give how much header you want)
func RequestDelete(client *http.Client, url string, data []byte, header url.Values) ([]byte, error) {
	return Request(client, http.MethodDelete, url, data, header)
}

// Description: Do an http.Request using http.Client
//
// Parameters:
//   - *http.Client
//   - Method(string) example: "GET", "POST" ..
//   - Url(string)
//   - data([]byte)
//   - header(...tools.Pair) (params-pack you can give how much header you want)
func Request(client *http.Client, method string, url string, data []byte, header url.Values) ([]byte, error) {
	//create request
	var body io.Reader = nil
	if data != nil {
		body = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(string(method), url, body)
	if err != nil {
		return nil, err
	}

	//fill request header
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	//	resp.StatusCode
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
