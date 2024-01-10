package main

import "fmt"

func getCombination(combination int, N int) [][]int {
	var result = make([][]int, 0)
	var comb = []int{}
	var level = combination

	var backtrack func(comb []int, level int, pos int, N int)
	backtrack = func(comb []int, level int, pos int, N int) {
		if level == 0 {
			temp := make([]int, combination)
			copy(temp, comb)
			result = append(result, temp)
		} else {
			for i := pos; i < N+1; i++ {
				comb = append(comb, i)
				backtrack(comb, level-1, i+1, N)
				comb = comb[:len(comb)-1]
			}
		}
	}
	backtrack(comb, level, 1, N)
	return result
}
func main() {
	N := 3
	combination := 2
	fmt.Println(getCombination(combination, N))
}
