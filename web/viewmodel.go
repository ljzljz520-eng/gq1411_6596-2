package web

import "miaoxiu.example/domain"

type ViewModel struct {
	Title    string
	Patterns []domain.Pattern
	Artworks []domain.Artwork
}

func MakeView(c domain.Catalog) ViewModel {
	return ViewModel{Title: "苗绣工艺专题站", Patterns: c.Patterns, Artworks: c.Artworks}
}
