package service

import "fmt"

type Handle struct{ closed bool }

func (h *Handle) Close() { h.closed = true }
func acquire() *Handle   { return &Handle{} }
func (s *RegistrationService) ProcessBatch(n int) error {
	active := 0
	for i := 0; i < n; i++ {
		h := acquire()
		defer h.Close()
		active++
		if active > 3 {
			return fmt.Errorf("resource quota exhausted")
		}
	}
	return nil
}
