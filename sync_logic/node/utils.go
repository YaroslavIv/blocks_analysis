package node

import (
	"sort"
	"sync_logic/db"
	"sync_logic/ram"
)

func getTopFive(r1, r2 db.ListRow) ram.Top {
	return getTop(r1, r2, 5)
}

func getTop(r1, r2 db.ListRow, n int) ram.Top {
	var out ram.Top
	r := append(r1, r2...)
	rm := make(map[string]int)

	for _, i := range r {
		rm[i.AddrFrom] += 1
		rm[i.AddrTo] += 1
	}

	keys := make([]string, 0, len(rm))

	for key := range rm {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return rm[keys[i]] > rm[keys[j]]
	})

	for i, k := range keys {
		if i == n {
			break
		}

		out = append(out, ram.TopAddr{Addr: k, Count: rm[k]})
	}

	return out
}
