package main

import (
	"fmt"
)

//func solveNQueens(n int) [][]string {
//
//}

func recurcive2(oneSolution [][]string) {
	for i := 0; i < len(oneSolution); i++ {
		for j := 0; j < len(oneSolution); j++ {
			flag := false
			if oneSolution[i][j] != "." && oneSolution[i][j] != "Q" {
				tempSolution := make([][]string, len(oneSolution))
				copy(tempSolution, oneSolution)
				fill(tempSolution, i, j)
				recurcive2(tempSolution)
				flag = true
			}
			if !flag {
				return
			}
		}
	}
	fmt.Println(oneSolution)
}

func fill(oneSolution [][]string, m, n int) {
	oneSolution[m][n] = "Q"
	for k := 0; k < len(oneSolution); k++ {
		if k != n {
			oneSolution[m][k] = "."
		}
	}
	for k := 0; k < len(oneSolution); k++ {
		if k != m {
			oneSolution[k][n] = "."
		}
	}
	for i, j := m, n; i+1 < len(oneSolution) && j+1 < len(oneSolution); {
		i++
		j++
		oneSolution[i][j] = "."
	}
	for i, j := m, n; i-1 > -1 && j-1 > -1; {
		i--
		j--
		oneSolution[i][j] = "."
	}
	for i, j := m, n; i+1 < len(oneSolution) && j-1 > -1; {
		i++
		j--
		oneSolution[i][j] = "."
	}
	for i, j := m, n; i-1 > -1 && j+1 < len(oneSolution); {
		i--
		j++
		oneSolution[i][j] = "."
	}
	fmt.Println(oneSolution)
}

func main() {
	one := make([][]string, 4)
	one[0] = make([]string, 4)
	one[1] = make([]string, 4)
	one[2] = make([]string, 4)
	one[3] = make([]string, 4)
	recurcive2(one)
}
