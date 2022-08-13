package stringutility

func Reversestring(t string) string {
	var s = ""
	for i := len(t)-1; i >= 0; i-- {
		s += string(t[i])
	}
	return s
}
