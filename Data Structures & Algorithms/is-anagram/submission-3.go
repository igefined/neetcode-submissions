func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := [26]uint8{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}

	return true
}
