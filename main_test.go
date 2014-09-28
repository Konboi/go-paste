package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	//"fmt"
)

type PingResult struct {
	Status  int `json:"status"`
	Results struct {
		Message string `json:"message"`
	} `json:"results"`
}

func TestPingHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(pingHandler))

	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Errorf("Someting Server Error: %s", err)
	}

	if res.StatusCode != 200 {
		t.Error("Status Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	pingResult := new(PingResult)
	err = json.Unmarshal(body, &pingResult)

	if err != nil {
		t.Errorf("JSON Parse Error: %s", err)
	}

	if pingResult.Status != 200 {
		t.Errorf("Ping Status is not 200. Status is %d", pingResult.Status)
	}

	if pingResult.Results.Message != "ok" {
		t.Errorf("Ping Results Message is not 'ok'. Result is %s", pingResult.Results.Message)
	}
}

func TestPostHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(postHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil && !strings.Contains(err.Error(), "Get /: stopped after 10 redirects") {
		t.Errorf("Someting Server Error: %s", err)
	}

	if res.StatusCode != 302 {
		t.Error("Do Not Redirect for Get Method")
	}

}
