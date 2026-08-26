package main

import (
	"miaoxiu.example/content"
	"testing"
)

func TestCatalogQueries(t *testing.T) {
	c := content.Seed()
	if len(c.Featured()) != 2 || len(c.Search("蝴蝶")) != 1 {
		t.Fatal("query")
	}
}
