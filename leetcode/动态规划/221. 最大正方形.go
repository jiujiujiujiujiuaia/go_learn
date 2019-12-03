package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalSquare([][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}))
}

//大问题拆成小问题，记忆化存储
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := prepare(matrix)
	best := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = help(matrix, dp, i, j)
			if dp[i][j] > best {
				best = dp[i][j]
			}
		}
	}
	return best * best
}

//优化动态规划，空间复杂度为0(n)
//func maximalSquare2(matrix [][]byte) int {
//	if len(matrix) == 0 {
//		return 0
//	}
//	dp := make([]int,len(matrix))
//	best := 0
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			if matrix[i][j] == '0' {
//				continue
//			}
//			dp[i][j] = help(matrix, dp, i, j)
//			if dp[i][j] > best {
//				best = dp[i][j]
//			}
//		}
//	}
//	return best * best
//}

func help(matrix [][]byte, dp [][]int, i, j int) int {
	if i == 0 || j == 0 {
		if matrix[i][j] == '0' {
			return 0
		} else {
			return 1
		}
	}
	min := minNum(minNum(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
	return min + 1
}

func prepare(matrix [][]byte) [][]int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	return dp
}

func minNum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
