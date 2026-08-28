package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"miaoxiu.example/domain"
	"path/filepath"
	"sync"
)

var buckets = []string{"patterns", "stitches", "artisans", "artworks", "registrations", "resources"}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(b)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error { return s.db.Close() }
func put[T any](s *Store, b, id string, v T) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(b)).Put([]byte(id), raw) })
}
func get[T any](s *Store, b, id string) (T, error) {
	var z T
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(b)).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("not found")
		}
		return json.Unmarshal(v, &z)
	})
	return z, e
}
func list[T any](s *Store, b string) ([]T, error) {
	out := []T{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(b)).ForEach(func(_, v []byte) error {
			var z T
			if e := json.Unmarshal(v, &z); e != nil {
				return e
			}
			out = append(out, z)
			return nil
		})
	})
	return out, e
}
func (s *Store) SavePattern(v domain.Pattern) error { return put(s, "patterns", v.ID, v) }
func (s *Store) SaveStitch(v domain.Stitch) error   { return put(s, "stitches", v.ID, v) }
func (s *Store) SaveArtisan(v domain.Artisan) error { return put(s, "artisans", v.ID, v) }
func (s *Store) SaveArtwork(v domain.Artwork) error { return put(s, "artworks", v.ID, v) }
func (s *Store) SaveRegistration(v domain.Registration) error {
	return put(s, "registrations", v.ID, v)
}
func (s *Store) SaveResource(v domain.Resource) error { return put(s, "resources", v.ID, v) }
func (s *Store) Pattern(id string) (domain.Pattern, error) {
	return get[domain.Pattern](s, "patterns", id)
}
func (s *Store) Registrations() ([]domain.Registration, error) {
	return list[domain.Registration](s, "registrations")
}
func (s *Store) Patterns() ([]domain.Pattern, error) { return list[domain.Pattern](s, "patterns") }
func (s *Store) Artworks() ([]domain.Artwork, error) { return list[domain.Artwork](s, "artworks") }
func (s *Store) Resource(id string) (domain.Resource, error) {
	return get[domain.Resource](s, "resources", id)
}
