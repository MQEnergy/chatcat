package chttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Request
// @Description: request
// @param method
// @param url
// @param data
// @return string
// @return error
// @author cx
func Request(method, url, data string) (string, error) {
	body := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	//req.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(response.Body)
	return string(b), err
}
