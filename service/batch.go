package service

import "fmt"

type Handle struct{ closed bool }

func (h *Handle) Close() { h.closed = true }
func acquire() *Handle   { return &Handle{} }
func (s *RegistrationService) ProcessBatch(n int) error {
	handles := []*Handle{}
	for i := 0; i < n; i++ {
		h := acquire()
		handles = append(handles, h)
		if len(handles) > 3 {
			return fmt.Errorf("resource quota exhausted")
		}
	}
	defer func() {
		for _, h := range handles {
			h.Close()
		}
	}()
	return nil
}
