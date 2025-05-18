package requests

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type Request interface {
	Send()
}

type GETRequest struct {
	Url string
}

func (r *GETRequest) Send() {
	resp, err := http.Get(r.Url)
	if err != nil {
		log.Printf("GET request to %s failed: %v", r.Url, err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}
	
	log.Printf("GET %s: %s", r.Url, string(body))
}

type POSTRequest struct {
	URL string
	Data []byte
}

func (r *POSTRequest) Send() {
	resp, err := http.Post(r.URL, "application/json", bytes.NewBuffer(r.Data))
	if err != nil {
		log.Printf("POST request to %s failed: %v", r.URL, err)
		return
	}
	defer resp.Body.Close()	
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}	
	
	log.Printf("POST %s: %s", r.URL, string(body))
}