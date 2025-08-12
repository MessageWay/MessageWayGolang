package MessageWay

func isValidHash(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '#' {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}
