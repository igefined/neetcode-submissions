func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		set[s[i]] += 1
	}

	for i := 0; i < len(t); i++ {
		el, ok := set[t[i]]
		if !ok {
			return false
		}

		if el > 1 {
			set[t[i]] -= 1
			continue
		}
		
		delete(set, t[i])
	}

	return true
}
