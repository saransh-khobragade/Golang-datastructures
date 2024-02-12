package main

import (
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func groupAnagrams(strs []string) [][]string {
	hmap := map[string][]string{}

	for _, v := range strs {
		str := sortString(v)

		if x, y := hmap[str]; y {
			x = append(x, v)
			hmap[str] = x
		} else {
			hmap[str] = []string{v}
		}
	}

	out := [][]string{}
	for _, v := range hmap {
		out = append(out, v)
	}
	return out
}
