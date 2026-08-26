package main

import (
	"miaoxiu.example/service"
	"miaoxiu.example/store"
	"os"
	"testing"
)

func TestRegistrationService(t *testing.T) {
	p := "svc.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	x := service.New(s)
	r, e := x.Register("甲", "123456", "周六上午")
	if e != nil || !r.Confirmed() {
		t.Fatal(e)
	}
}
