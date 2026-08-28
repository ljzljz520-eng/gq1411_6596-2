package service

import "miaoxiu.example/domain"

func StatusCounts(rs []domain.Registration) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func Confirmed(rs []domain.Registration) []domain.Registration {
	out := []domain.Registration{}
	for _, r := range rs {
		if r.Confirmed() {
			out = append(out, r)
		}
	}
	return out
}
