package analytics

import (
	"miaoxiu.example/domain"
	"sort"
)

type Summary struct{ PatternCount, ArtworkCount, ActiveArtisanCount int }

func Build(c domain.Catalog) Summary {
	n := 0
	for _, a := range c.Artisans {
		if a.Active {
			n++
		}
	}
	return Summary{len(c.Patterns), len(c.Artworks), n}
}
func RankPatterns(ps []domain.Pattern) []domain.Pattern {
	out := append([]domain.Pattern{}, ps...)
	sort.Slice(out, func(i, j int) bool { return out[i].Featured && !out[j].Featured })
	return out
}
