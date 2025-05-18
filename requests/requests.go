package requests

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Err        error
}

type Request interface {
	Send() (*Response, error)
}

type GETRequest struct {
	Url string
}

func (r *GETRequest) Send() (*Response, error) {
	resp, err := http.Get(r.Url)
	if err != nil {
		log.Printf("GET request to %s failed: %v", r.Url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	log.Printf("GET %s: %s", r.Url, string(body))
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
		Err:        nil,
	}, nil
}

type POSTRequest struct {
	URL  string
	Data []byte
}

func (r *POSTRequest) Send() (*Response, error) {
	resp, err := http.Post(r.URL, "application/json", bytes.NewBuffer(r.Data))
	if err != nil {
		log.Printf("POST request to %s failed: %v", r.URL, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil, err
	}

	log.Printf("POST %s: %s", r.URL, string(body))
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
		Err:        nil,
	}, nil
}
