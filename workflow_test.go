package main

import (
	"miaoxiu.example/content"
	"miaoxiu.example/service"
	"miaoxiu.example/store"
	"os"
	"testing"
)

func TestWorkflowBrowse(t *testing.T) {
	c := content.Seed()
	if len(c.Patterns) == 0 || len(c.Artworks) == 0 {
		t.Fatal("empty site")
	}
}
func TestWorkflowRegistration(t *testing.T) {
	p := "wf.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	if _, e := service.New(s).Register("访客", "1380000", "周日下午"); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain44(t *testing.T) {
	p := "bug.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	if e := service.New(s).ProcessBatch(4); e == nil {
		t.Fatal("expected resource quota exhaustion from batch")
	}
}
