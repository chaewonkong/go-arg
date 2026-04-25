// Package anagram
//
//   - https://leetcode.com/problems/find-all-anagrams-in-a-string/
//   - Time: O(n)
//   - Space: O(1)
package anagram

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	result := []int{}
	var pCount, wCount [26]int

	for _, c := range p {
		pCount[c-'a']++
	}

	for i := range len(s) {
		wCount[s[i]-'a']++ // mark

		if i >= len(p) {
			first := s[i-len(p)] - 'a'
			wCount[first]--
		}

		if i >= len(p)-1 && pCount == wCount {
			result = append(result, i-len(p)+1)
		}

	}

	return result
}
