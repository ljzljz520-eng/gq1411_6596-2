package main

import (
	"miaoxiu.example/domain"
	"miaoxiu.example/store"
	"os"
	"testing"
)

func TestStoreEntities(t *testing.T) {
	p := "test.db"
	defer os.Remove(p)
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.SavePattern(domain.Pattern{ID: "x", Name: "纹", Meaning: "意"}); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Pattern("x"); e != nil {
		t.Fatal(e)
	}
}
