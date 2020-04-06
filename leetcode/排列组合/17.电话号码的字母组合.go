package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

//2020/2/27
var res []string
var maps = map[int][]string{2: {"a", "b", "c"}, 3: {"d", "e", "f"}, 4: {"g", "h", "i"}, 5: {"j", "k", "l"}, 6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"}, 8: {"t", "u", "v"}, 9: {"w", "x", "y", "z"}}

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if digits != "" {
		help(digits, "")
	}
	return res
}

func help(digits string, str string) {
	if digits == "" {
		res = append(res, str)
		return
	}

	for _, v := range maps[int(digits[0])-48] {
		help(digits[1:], str+v)
	}
}

// @lc code=end
