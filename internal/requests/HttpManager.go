package requests

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPManager stuff
type HTTPManager struct {
	client *http.Client
}

// HTTPManagerResponse stuff
type HTTPManagerResponse struct {
	responseCode int
	responseData string
	headers      map[string]string
}

// GetResponseCode stuff
func (r *HTTPManagerResponse) GetResponseCode() int {
	return r.responseCode
}

// GetResponseData stuff
func (r *HTTPManagerResponse) GetResponseData() string {
	return r.responseData
}

// GetHeaders stuff
func (r *HTTPManagerResponse) GetHeaders() map[string]string {
	return r.headers
}

// CreateHTTPManagerResponse stuff
func CreateHTTPManagerResponse(resp *http.Response) (*HTTPManagerResponse, error) {

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	for k, v := range resp.Header {
		headers[k] = v[0] // todo: broken?
	}

	r := &HTTPManagerResponse{responseCode: resp.StatusCode, responseData: string(body), headers: headers}
	return r, nil
}

// CreateHTTPManager stuff
func CreateHTTPManager() *HTTPManager {
	client := &http.Client{}
	mgr := &HTTPManager{client}
	return mgr
}

func (mgr *HTTPManager) performRequest(req *http.Request, requestHeaders map[string]string) (*HTTPManagerResponse, error) {
	for k, v := range requestHeaders {
		req.Header.Add(k, v)
	}

	resp, err := mgr.client.Do(req)
	if err != nil {
		return nil, err
	}

	return CreateHTTPManagerResponse(resp)
}

// Get stuff
func (mgr *HTTPManager) Get(url string, requestHeaders map[string]string) (*HTTPManagerResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return mgr.performRequest(req, requestHeaders)
}

// Post stuff
func (mgr *HTTPManager) Post(url string, body string, requestHeaders map[string]string) (*HTTPManagerResponse, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	return mgr.performRequest(req, requestHeaders)
}
