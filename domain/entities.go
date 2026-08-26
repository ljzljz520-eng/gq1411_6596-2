package domain

import "time"

type Pattern struct {
	ID, Name, Meaning, Region, Image string
	Featured                         bool
}
type Stitch struct{ ID, Name, Technique, Difficulty, Description string }
type Artisan struct {
	ID, Name, Village, Bio, Portrait string
	Active                           bool
}
type Artwork struct {
	ID, Title, PatternID, ArtisanID, Image, Story string
	Year                                          int
}
type Registration struct {
	ID, Name, Phone, Session, Status string
	CreatedAt                        time.Time
}
type Resource struct {
	ID          string
	Limit, Used int
}

func (p Pattern) Valid() bool      { return p.ID != "" && p.Name != "" && p.Meaning != "" }
func (s Stitch) Valid() bool       { return s.ID != "" && s.Name != "" }
func (a Artisan) Valid() bool      { return a.ID != "" && a.Name != "" }
func (a Artwork) Valid() bool      { return a.ID != "" && a.Title != "" }
func (r Registration) Valid() bool { return r.ID != "" && r.Name != "" && r.Phone != "" }
func (r Resource) Available() bool { return r.Used < r.Limit }
