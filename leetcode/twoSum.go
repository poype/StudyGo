package main

func twoSum(nums []int, target int) []int {
	numIdxMap := make(map[int]int)

	for i, val := range nums {
		if idx, ok := numIdxMap[val]; ok {
			return []int{idx, i}
		}
		numIdxMap[target-val] = i
	}
	return nil
}

//var nums = []int{2, 7, 11, 15}
//var target = 9

var nums = []int{3, 3}
var target = 6
