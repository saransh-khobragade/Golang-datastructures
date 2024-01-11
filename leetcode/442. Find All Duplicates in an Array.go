package leetcode

func findDuplicates(nums []int) []int {
	hashMap := make(map[int]int)
	result := []int{}
	for _, v := range nums {
		if _, x := hashMap[v]; x {
			result = append(result, v)
		} else {
			hashMap[v] = 1
		}
	}
	return result
}
