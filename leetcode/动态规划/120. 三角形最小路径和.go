package main

import "fmt"

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := len(triangle) - 1; i > -1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i+1 < len(triangle) {
				dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
			} else {
				dp[i][j] = triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
