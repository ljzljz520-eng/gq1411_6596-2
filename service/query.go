package service

import (
	"miaoxiu.example/domain"
	"miaoxiu.example/store"
	"strings"
)

type QueryService struct {
	Store   *store.Store
	Catalog domain.Catalog
}

func (q *QueryService) FindPatterns(term string) []domain.Pattern {
	out := []domain.Pattern{}
	for _, p := range q.Catalog.Patterns {
		if term == "" || strings.Contains(p.Name, term) || strings.Contains(p.Meaning, term) {
			out = append(out, p)
		}
	}
	return out
}
func (q *QueryService) Gallery() []domain.Artwork {
	a, _ := q.Store.Artworks()
	if len(a) == 0 {
		return q.Catalog.Artworks
	}
	return a
}
