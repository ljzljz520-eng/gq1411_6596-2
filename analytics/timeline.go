package analytics

import "miaoxiu.example/domain"

func Years(as []domain.Artwork) []int {
	out := []int{}
	seen := map[int]bool{}
	for _, a := range as {
		if !seen[a.Year] {
			seen[a.Year] = true
			out = append(out, a.Year)
		}
	}
	return out
}
func Latest(as []domain.Artwork) domain.Artwork {
	var z domain.Artwork
	for _, a := range as {
		if a.Year > z.Year {
			z = a
		}
	}
	return z
}
