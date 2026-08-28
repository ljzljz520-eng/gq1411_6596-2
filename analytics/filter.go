package analytics

import "miaoxiu.example/domain"

func ByRegion(ps []domain.Pattern, region string) []domain.Pattern {
	out := []domain.Pattern{}
	for _, p := range ps {
		if region == "" || p.Region == region {
			out = append(out, p)
		}
	}
	return out
}
func Meanings(ps []domain.Pattern) map[string]int {
	m := map[string]int{}
	for _, p := range ps {
		m[p.Meaning]++
	}
	return m
}
