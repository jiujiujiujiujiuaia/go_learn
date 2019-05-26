package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}

func maxArea(height []int) int {
	return twoPoint(height)
}

//复杂度为o(n^2)
func complexity_n2(height []int) int {
	best := 0
	for i, val := range height {
		witdh := 0
		for j, v := range height {
			if val <= v && math.Abs(float64(i)-float64(j)) > float64(witdh) {
				witdh = int(math.Abs(float64(i) - float64(j)))
			}
		}
		if witdh*val > best {
			best = witdh * val
		}
	}
	return best
}

//双指针 复杂度为o(n)
func twoPoint(height []int) int {
	best := 0
	i := 0
	for j := len(height) - 1; i < j; {
		if best < (j-i)*If(height[i], height[j]) {
			best = (j - i) * If(height[i], height[j])
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return best
}

func If(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
