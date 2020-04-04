package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

//解法一：动态规划的解法
//dp[i]表示的是以s[i]结尾的最长子串合法括号
//s[i]="(" dp[i] = 0
//s[i]=")" dp[i] = dp[i-2]+2 or ...

//case0 ()
//case1 (())
//case2 ()(())

//如果使用了动态规划做，如果面试官要求你求输出这个子串怎么办？
//再加一个状态数组，表示的是，如果以i结尾，可以构成合法括号，则记录最前一个左括号的索引
func longestValidParentheses1(s string) int {
	dp := make([]int, len(s)+1)
	maxLength := 0
	//注意dp索引的定义和s索引的定义不同
	for i := 2; i <= len(s); i++ {
		if s[i-1] == ')' {
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else if s[i-2] == ')' {
				if i-dp[i-1]-2 >= 0 && s[i-dp[i-1]-2] == '(' && dp[i-1] > 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				}
			}
		}

		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}
	return maxLength
}

//解法二：使用栈保存索引来做，这样不仅可以记录最长的合法括号，还可以输出出来哪个是最长的
// func longestValidParentheses(s string) int {
// 	int maxans = 0;
// 	Stack<Integer> stack = new Stack<>();
// 	stack.push(-1);
// 	for (int i = 0; i < s.length(); i++) {
// 		if (s.charAt(i) == '(') {
// 			stack.push(i);
// 		} else {
// 			stack.pop();
//点睛之笔，关于为什么这里为空的时候要把当前push进去
//因为)()时，pop)然后由于stack为空了，栈中没有了索引坐标作为锚定，
//因此把)push入作为坐标的锚定
// 			if (stack.empty()) {
// 				stack.push(i);
// 			} else {
// 				maxans = Math.max(maxans, i - stack.peek());
// 			}
// 		}
// 	}
// 	return maxans;
// }

//解法三，分别从左到右再从右到左遍历，为什么要遍历两遍呢？
//因为从左到右无法解决(() 也就是左边括号多的情况，left就是比right大，但是又无法解决
//同理从右到左也无法解决())情况

// @lc code=end
