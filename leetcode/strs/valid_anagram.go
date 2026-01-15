package strs

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}

	for i := range len(s) {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}
