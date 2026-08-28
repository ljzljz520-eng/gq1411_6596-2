package service

import (
	"fmt"
	"miaoxiu.example/domain"
	"miaoxiu.example/store"
	"strings"
)

type RegistrationService struct{ Store *store.Store }

func New(s *store.Store) *RegistrationService { return &RegistrationService{Store: s} }
func (s *RegistrationService) Register(name, phone, session string) (domain.Registration, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(phone) == "" {
		return domain.Registration{}, fmt.Errorf("name and phone required")
	}
	r := domain.Registration{ID: fmt.Sprintf("reg-%d", s.next()), Name: name, Phone: phone, Session: session, Status: "confirmed"}
	if e := s.Store.SaveRegistration(r); e != nil {
		return r, e
	}
	return r, nil
}
func (s *RegistrationService) next() int                   { rs, _ := s.Store.Registrations(); return len(rs) + 1 }
func (s *RegistrationService) List() []domain.Registration { r, _ := s.Store.Registrations(); return r }
func (s *RegistrationService) SeedResource() error {
	return s.Store.SaveResource(domain.Resource{ID: "experience", Limit: 3, Used: 0})
}
func (s *RegistrationService) ConsumeResource() error {
	r, e := s.Store.Resource("experience")
	if e != nil {
		return e
	}
	if !r.Available() {
		return fmt.Errorf("quota exhausted")
	}
	r.Used++
	return s.Store.SaveResource(r)
}
