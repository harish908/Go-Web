package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 5 * time.Second,
		},
	}
}

func TestGetIdeas(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/api/ideas", bytes.NewReader(nil))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(body)
}
