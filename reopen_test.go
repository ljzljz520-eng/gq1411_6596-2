package main

import (
	"miaoxiu.example/domain"
	"miaoxiu.example/store"
	"os"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "reopen.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	s.SaveRegistration(domain.Registration{ID: "r1", Name: "甲", Phone: "123456", Status: "confirmed"})
	s.Close()
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	rs, _ := s.Registrations()
	if len(rs) != 1 {
		t.Fatal("lost")
	}
}
