package main

import (
	"fmt"
)

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	step := n - 1
	loop := n / 2
	opp := 0
	count := 1
	for ; loop > 0; loop-- {
		for i := opp; i < step+opp; i++ {
			matrix[opp][i] = count
			count++
		} //up

		for i := opp; i < step+opp; i++ {
			matrix[i][n-1-opp] = count

			count++
		} //right

		for i := step + opp; i > opp; i-- {
			matrix[n-1-opp][i] = count
			count++
		} //bottom

		for i := step + opp; i > opp; i-- {
			matrix[i][opp] = count
			count++
		} //left
		opp++
		step -= 2
	}
	if n%2 == 1 {
		matrix[n/2][n/2] = count
	}
	return matrix
}

func main() {
	m := generateMatrix(7)
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
