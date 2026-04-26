// Package longestrepl
//
//   - https://leetcode.com/problems/longest-repeating-character-replacement/description/
//   - time: O(N)
//   - space: O(1)
package longestrepl

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	maxFreq := 0
	left := 0
	result := 0

	for right := range len(s) {
		freq[s[right]-'A']++ // freq에 기록
		maxFreq = max(maxFreq, freq[s[right]-'A'])

		// shrink: window - maxFreq > k
		if (right-left+1)-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}
