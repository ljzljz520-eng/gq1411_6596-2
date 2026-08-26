package domain

type Catalog struct {
	Patterns []Pattern
	Stitches []Stitch
	Artisans []Artisan
	Artworks []Artwork
}

func (c Catalog) Pattern(id string) (Pattern, bool) {
	for _, p := range c.Patterns {
		if p.ID == id {
			return p, true
		}
	}
	return Pattern{}, false
}
func (c Catalog) Artisan(id string) (Artisan, bool) {
	for _, a := range c.Artisans {
		if a.ID == id {
			return a, true
		}
	}
	return Artisan{}, false
}
func (c Catalog) Artwork(id string) (Artwork, bool) {
	for _, a := range c.Artworks {
		if a.ID == id {
			return a, true
		}
	}
	return Artwork{}, false
}
func (c Catalog) Featured() []Pattern {
	out := []Pattern{}
	for _, p := range c.Patterns {
		if p.Featured {
			out = append(out, p)
		}
	}
	return out
}
func (c Catalog) Search(q string) []Pattern {
	out := []Pattern{}
	for _, p := range c.Patterns {
		if q == "" || contains(p.Name, q) || contains(p.Meaning, q) {
			out = append(out, p)
		}
	}
	return out
}
func contains(a, b string) bool {
	if b == "" {
		return true
	}
	for i := 0; i+len(b) <= len(a); i++ {
		if a[i:i+len(b)] == b {
			return true
		}
	}
	return false
}
