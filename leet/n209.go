package main

import "log"

//209. 长度最小的子数组
//中等
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的
//子数组
//[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//示例 2：
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//示例 3：
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//提示：
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 104

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	i := 0
	result := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			sub_result := j - i + 1
			if sub_result < result {
				result = sub_result
			}
			sum -= nums[i]
			i += 1
		}
	}
	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}

}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	log.Println(minSubArrayLen(7, nums))
}
