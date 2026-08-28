package registration

func Slots() []string {
	return []string{"周六上午", "周六下午", "周日上午", "周日下午"}
}
func SlotIndex(v string) int {
	for i, s := range Slots() {
		if s == v {
			return i
		}
	}
	return -1
}
