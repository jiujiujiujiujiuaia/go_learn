package main

import (
	"fmt"
	"strings"
)

//最大值,math库中有这些常量！！！math.MaxInt8
const INT_MAX = int(int32(^uint32(0) >> 1))
const INT_MIN = int(-int32(^uint32(0)>>1) - 1)

func main() {
	str := "9223372036854775808"
	fmt.Println(myAtoi(str))
}

func myAtoi(str string) int {
	j := 0
	i := -1
	flag := true
	res := 0
	str = strings.Trim(str, " ")

	if len(str) >= 1 && (str[0] > 57 || str[0] < 48) {
		if str[0] == 43 || str[0] == 45 {
			if len(str) < 2 || str[1] > 57 || str[1] < 48 {
				return 0
			}
		} else {
			return 0
		}
	}

	for index, val := range str {
		if val <= 57 && val >= 48 {
			if i == -1 {
				i = index
				j = index
			} else if index-1 == j {
				j = index
			}

		}
	}
	if i == -1 {
		return 0
	}

	if i-1 >= 0 && str[i-1] == 45 {
		flag = false
	}

	for m := 1; j >= i; j-- {
		res = (int(str[j])-48)*m + res
		m *= 10
	}

	if !flag {
		res = -res
	}

	if res >= INT_MAX {
		return INT_MAX
	} else if res <= INT_MIN {
		return INT_MIN
	} else {
		return res
	}

}
