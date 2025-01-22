package main

import (
	"fmt"
)

// 好未来随机笔试
func maxAscendingSum(nums []int) int {
	var lower []int
	sum := 0
	start := nums[0]
	lower = append(lower, nums[0])
	for i := 1; i < len(nums); i++ {
		if start > nums[i] {
			start = nums[i]
			lower = append(lower, nums[i])
		} else {
			start = nums[i]
			tmp := 0
			for l := range lower {
				tmp += lower[l]
			}
			if tmp > sum {
				sum = tmp
			}
			lower = []int{nums[i]}
		}
	}
	return sum
}

// 好未来随机笔试
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	arr1 := []int{50, 10, 20, 30, 5, 10}
	arr2 := []int{2, 7, 11, 15}

	fmt.Println(maxAscendingSum(arr1))
	fmt.Println(twoSum(arr2, 9))
}
