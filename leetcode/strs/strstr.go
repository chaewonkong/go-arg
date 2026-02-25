package strs

func strStr(haystack string, needle string) int {
	for i := range len(haystack) {
		if haystack[i] == needle[0] {
			found := true
			for j := 0; j < len(needle); j++ {
				if i+j >= len(haystack) {
					found = false
					break
				}
				if haystack[i+j] != needle[j] {
					found = false
					break
				}
			}

			if found {
				return i
			}
		}
	}

	return -1
}
