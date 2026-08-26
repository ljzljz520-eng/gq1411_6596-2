package main

import (
	"log"
	"miaoxiu.example/content"
	"miaoxiu.example/service"
	"miaoxiu.example/store"
	"miaoxiu.example/web"
	"net/http"
	"os"
)

func main() {
	path := "miaoxiu.db"
	if v := os.Getenv("MIAOXIU_DB"); v != "" {
		path = v
	}
	st, e := store.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	c := content.Seed()
	for _, p := range c.Patterns {
		st.SavePattern(p)
	}
	svc := service.New(st)
	svc.SeedResource()
	log.Println(http.ListenAndServe(":8080", web.New(c, svc).Routes()))
}
