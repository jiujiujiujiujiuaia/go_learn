package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start

//边界条件（1）为空（2）本身不构成一个完整的ip地址，s小于4或大于16
//（3）"010010" -> "01.0.01.0" 01不成立但是api会把01解析成1
//（4）考虑题目给的数据中是否只有数字？询问面试官

var result []string

func restoreIpAddresses(s string) []string {
	result = make([]string, 0)
	help(s, 0, "", 0)
	return result
}

func help(s string, start int, res string, cnt int) {
	//(1)递归终止条件，可以参考上面的边界情况
	if cnt == 4 && start != len(s) {
		return
	} else if cnt == 4 && start == len(s) {
		result = append(result, res[0:len(res)-1])
		fmt.Println(result)
		return
	}

	for i := start + 1; i <= len(s); i++ {
		if val, _ := strconv.Atoi(s[start:i]); val <= 255 {
			//（2）排除掉001 00 这种情况 api没办法识别
			if len(s[start:i]) > 1 && s[start] == '0' {
				continue
			}

			//（3）这里其实有一点值得注意，由于还会回溯,所以不能执行下面注释的操作
			//res+ = s[start:i]+"."
			//cnt ++
			//因为递归完后还会回到这里，循环执行下一个起点，这样会影响另一个回溯开始的头
			help(s, i, res+s[start:i]+".", cnt+1)
		}
	}
}

// func main() {
// 	s := "001"
// 	val, _ := strconv.ParseInt(s, 10, 64)
// 	fmt.Println(val)
// }

// @lc code=end
