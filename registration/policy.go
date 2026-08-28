package registration

import "strings"

func NormalizePhone(v string) string {
	return strings.ReplaceAll(strings.ReplaceAll(v, " ", ""), "-", "")
}
func SessionAllowed(v string) bool {
	return v == "周六上午" || v == "周日下午" || v == "周日上午"
}
func ValidContact(name, phone string) bool {
	return strings.TrimSpace(name) != "" && len(NormalizePhone(phone)) >= 6
}
