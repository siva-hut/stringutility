package stringutility

func Reversestring(t string) string {
	var s = ""
	for i := 0; i < len(t); i++ {
		s += string(t[i])
	}
	return s
}
