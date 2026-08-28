package domain

import "fmt"

func (p Pattern) Label() string        { return fmt.Sprintf("%s · %s", p.Name, p.Meaning) }
func (a Artisan) Label() string        { return fmt.Sprintf("%s（%s）", a.Name, a.Village) }
func (w Artwork) Label() string        { return fmt.Sprintf("%s %d", w.Title, w.Year) }
func (r Registration) Confirmed() bool { return r.Status == "confirmed" }
