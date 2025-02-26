package main

import (
	"fmt"
)

//349. 两个数组的交集
//简单

// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
//
// 提示：
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
func intersection(nums1 []int, nums2 []int) []int {
	// 哈希表用法
	n1, n2 := make([]int, 1001), make([]int, 1001)
	var n_new []int
	for _, n := range nums1 {
		n1[n] = 1
	}
	for _, n := range nums2 {
		n2[n] = 1
	}
	for i, n := range n1 {
		if n == 1 && n1[i] == n2[i] {
			n_new = append(n_new, i)
		}
	}
	return n_new
}
func main() {
	nums1 := []int{1, 2, 5}
	nums2 := []int{3, 4, 5}
	fmt.Println(intersection(nums1, nums2))
}
