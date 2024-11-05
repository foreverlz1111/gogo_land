package main

import "log"

// 冒排 for两次，第二层每次都减少
func Buble(myarray []int) []int {
	for i := 0; i < len(myarray); i++ {
		for j := i + 1; j < len(myarray); j++ {
			if myarray[i] > myarray[j] {
				myarray[i], myarray[j] = myarray[j], myarray[i]
			}
		}
	}
	return myarray
}

// 插入排序，遍历每个，最小的放入前一列；重复
func Inseting(myarray []int) []int {
	// 从第头开始
	for i := range myarray {
		last_index := i - 1
		current := myarray[i]
		for last_index >= 0 && myarray[last_index] > current {
			myarray[last_index+1] = myarray[last_index]
			last_index -= 1
		}
		myarray[last_index+1] = current
	}
	return myarray
}

// 快速排序的最坏运行情况是 O(n²)，比如说顺序数列的快排。但它的平摊期望时间是 O(nlogn)，且 O(nlogn) 记号中隐含的常数因子很小，比复杂度稳定等于 O(nlogn) 的归并排序要小很多。所以，对绝大多数顺序性较弱的随机数列而言，快速排序总是优于归并排序。
func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 选择排序就是每一次都选择最大（最小）的放到最后（最前）
func Selection(myarray []int) []int {
	for i := 0; i < len(myarray)-1; i++ {
		minit := i
		for j := i + 1; j < len(myarray); j++ {
			if myarray[j] < myarray[minit] {
				minit = j
			}
		}
		myarray[i], myarray[minit] = myarray[minit], myarray[i]
	}
	return myarray
}
func main() {
	myarray := []int{22, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	log.Println(Selection(myarray))

}
