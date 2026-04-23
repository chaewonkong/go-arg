// Package accountsmerge
//
//   - https://leetcode.com/problems/accounts-merge/description/
//   - time: O(N·α(N))
//   - space: O(N)
package accountsmerge

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	emailToName := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py { // 한 account 내에서 연결
			parent[px] = py // compression
		}
	}

	for _, acc := range accounts {
		name := acc[0]
		for _, email := range acc[1:] {
			if _, ok := parent[email]; !ok {
				parent[email] = email // 자기 자신이 root
			}
			emailToName[email] = name

			// 첫 email 기준으로 union
			union(acc[1], email)
		}
	}

	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}

	results := [][]string{}
	for root, emails := range groups {
		sort.Strings(emails)
		accunt := append([]string{emailToName[root]}, emails...)
		results = append(results, accunt)
	}

	return results
}
