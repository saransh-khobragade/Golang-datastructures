package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	hmap := map[string]int{}

	if len(s) != len(t) {
		return false
	}
	for _, v := range s {
		if _, x := hmap[string(v)]; x {
			hmap[string(v)] += 1
		} else {
			hmap[string(v)] = 1
		}
	}
	for _, v := range t {
		if _, x := hmap[string(v)]; x {
			hmap[string(v)] -= 1
		}
	}
	for _, v := range hmap {
		if v > 0 {
			return false
		}
	}
	return true
}
