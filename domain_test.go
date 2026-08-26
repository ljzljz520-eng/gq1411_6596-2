package main

import (
	"miaoxiu.example/content"
	"testing"
)

func TestCatalogSeed(t *testing.T) {
	if !content.Validate(content.Seed()) {
		t.Fatal("invalid seed")
	}
}
