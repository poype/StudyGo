package main

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		moveToNextZero(&i, nums)
		moveToNextNonZero(&j, nums)

		if j < i {
			j++
			continue
		}

		if i < len(nums) && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func moveToNextZero(idx *int, nums []int) {
	i := *idx
	for i < len(nums) && nums[i] != 0 {
		i++
	}
	*idx = i
}

func moveToNextNonZero(idx *int, nums []int) {
	i := *idx
	for i < len(nums) && nums[i] == 0 {
		i++
	}
	*idx = i
}
