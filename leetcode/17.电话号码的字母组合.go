package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (51.59%)
 * Likes:    518
 * Dislikes: 0
 * Total Accepted:    63.7K
 * Total Submissions: 123.3K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start

var maps = map[int][]string{2: {"a", "b", "c"}, 3: {"d", "e", "f"}, 4: {"g", "h", "i"}, 5: {"j", "k", "l"}, 6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"}, 8: {"t", "u", "v"}, 9: {"w", "x", "y", "z"}}

var res = make([]string, 0)

func letterCombinations(digits string) []string {
	if digits != "" {
		help(digits, "")
	}
	return res
}

func help(str string, com string) {
	if str == "" {
		res = append(res, com)
		return
	}

	for _, v := range maps[int(str[0])-48] {
		help(str[1:], com+v)
	}
}

// @lc code=end
