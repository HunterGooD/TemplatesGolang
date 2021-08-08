package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingStruct struct {
	Message string `json:"message"`
}

func TestPing(t *testing.T) {
	app := App{}

	ts := httptest.NewServer(app.NewServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	pong := PingStruct{}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error on read data body")
	}

	if err := json.Unmarshal(buf, &pong); err != nil {
		t.Fatalf("Error on parse data %s", buf)
	}

	if pong.Message != "pong" {
		t.Fatalf("Error answer server expected \"pong\", got '%s'", pong.Message)
	}
}
