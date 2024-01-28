package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	result := []int{}
	str := ""
	for i := 0; i < len(num); i++ {
		str += strconv.Itoa(num[i])
	}

	j := new(big.Int)
	j.SetString(str, 10)
	j.Add(j, big.NewInt(int64(k)))

	d := big.NewInt(10)
	for j.Cmp(big.NewInt(0)) == 1 {
		mod := new(big.Int)
		mod.Mod(j, d)
		result = append(result, int(mod.Int64()))
		j.Div(j, d)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	num := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	k := 516

	fmt.Println(addToArrayForm(num, k))
}
