package validpalindrome

func isPalindrome(s string) bool {
	// remove non alphanumeric characters
	removed := ""
	for _, c := range s {
		if c >= 48 && c <= 57 { // number
			removed += string(c)
			continue
		}
		if c >= 65 && c <= 90 { // uppercase
			removed += string(c + 32) // to lower case
			continue
		}

		if c >= 97 && c <= 122 { // lowercase
			removed += string(c)
			continue
		}
	}

	length := len(removed)
	for i, c := range removed {
		j := length - i - 1
		if i >= j {
			break
		}

		if c != rune(removed[j]) {
			return false
		}
	}
	return true
}
