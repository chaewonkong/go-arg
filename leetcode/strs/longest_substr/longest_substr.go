package substr

// time complexity: O(n^2); space complexity: O(min(n, charset))
func lengthOfLongestSubstring(s string) int {
	maxL := 0
	for i := range len(s) {
		dup := make(map[byte]struct{})
		dup[s[i]] = struct{}{}
		cnt := 1
		for j := i + 1; j < len(s); j++ {

			if _, ok := dup[s[j]]; ok {
				break
			}
			dup[s[j]] = struct{}{}
			cnt++
		}
		maxL = max(maxL, cnt)
	}

	return maxL
}

// time complexity: O(n); space complexity: O(min(n, charset))
func lengthOfLongestSubstringWithSlidingWindow(s string) int {
	left := 0
	maxL := 0
	lastSeen := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		maxL = max(maxL, right-left+1)
	}

	return maxL
}
