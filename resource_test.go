package main

import (
	"miaoxiu.example/service"
	"miaoxiu.example/store"
	"os"
	"testing"
)

func TestBatchQuota(t *testing.T) {
	p := "batch.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	x := service.New(s)
	if e := x.ProcessBatch(4); e == nil {
		t.Fatal("quota should be reported")
	}
}
