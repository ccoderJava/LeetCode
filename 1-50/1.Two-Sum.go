package fun50

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		if ok {
			return []int{targetNumIndex, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}

func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	if nums[res[0]]+nums[res[1]] != target {
		fmt.Printf("twoSum(%v,%v) returned %d, want %d", nums, target, res, 3)
	}
}
