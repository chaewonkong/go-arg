package sort

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	start := 1
	end := n
	mid := (start + end) / 2

	for start < end {
		mid = (start + end) / 2

		if isBadVersion(mid) {
			end = mid
			continue
		}

		start = mid + 1
	}

	return start // if start == end
}
