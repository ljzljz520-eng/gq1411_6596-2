package main

import (
	"miaoxiu.example/content"
	"miaoxiu.example/web"
	"net/http/httptest"
	"testing"
)

func TestWebHome(t *testing.T) {
	r := httptest.NewRecorder()
	web.New(content.Seed(), nil).Routes().ServeHTTP(r, httptest.NewRequest("GET", "/", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
