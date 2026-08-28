package content

import "miaoxiu.example/domain"

func Merge(base domain.Catalog, extra domain.Catalog) domain.Catalog {
	out := base
	out.Patterns = append(out.Patterns, extra.Patterns...)
	out.Stitches = append(out.Stitches, extra.Stitches...)
	out.Artisans = append(out.Artisans, extra.Artisans...)
	out.Artworks = append(out.Artworks, extra.Artworks...)
	return out
}
func Validate(c domain.Catalog) bool {
	for _, p := range c.Patterns {
		if !p.Valid() {
			return false
		}
	}
	for _, s := range c.Stitches {
		if !s.Valid() {
			return false
		}
	}
	return true
}
