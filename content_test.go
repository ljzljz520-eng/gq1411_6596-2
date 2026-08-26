package main

import (
	"miaoxiu.example/content"
	"miaoxiu.example/domain"
	"testing"
)

func TestMergeContent(t *testing.T) {
	c := content.Merge(content.Seed(), domain.Catalog{})
	if len(c.Patterns) != 4 {
		t.Fatal("merge")
	}
}
