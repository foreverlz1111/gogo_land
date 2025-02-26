package main

import "log"

//4. 寻找两个正序数组的中位数
//困难
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 因为有序，先确定最小的停止-进行双指针填充，再补充剩余的信息到总的数组，得到总数组
	l1, l2 := 0, 0
	var sum []int
	var minlen int
	if len(nums1) >= len(nums2) {
		log.Println(len(nums1), len(nums2))
		minlen = len(nums2)
	} else {
		minlen = len(nums1)
	}

	for l1 < minlen && l2 < minlen {
		if nums1[l1] < nums2[l2] {
			//log.Println("nums1[l1] < nums2[l2]")
			sum = append(sum, nums1[l1])
			l1++
		} else if nums1[l1] > nums2[l2] {
			//log.Println("nums1[l1] > nums2[l2]")
			sum = append(sum, nums2[l2])
			l2++
		} else if nums1[l1] == nums2[l2] {
			//log.Println("nums1[l1] = nums2[l2]")
			sum = append(sum, nums1[l1])
			sum = append(sum, nums2[l2])
			l1++
			l2++
		}
	}

	if l1 == minlen {
		sum = append(sum, nums2[l2:]...)
	} else {
		sum = append(sum, nums1[l1:]...)
	}
	log.Println(sum)
	if len(sum)%2 == 1 {
		return float64(sum[len(sum)/2])
	} else {

		mid_sum := sum[len(sum)/2-1] + sum[len(sum)/2]
		return float64(mid_sum) / 2.00000
	}

}
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	log.Println(findMedianSortedArrays(nums1, nums2))
}
