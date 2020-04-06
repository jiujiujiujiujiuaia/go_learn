package main

/*
 * @lc app=leetcode.cn id=151 lang=golang
 * [151] 翻转字符串里的单词

 */

// @lc code=start

//2020/3/15 像这种以某个符号作为滑动窗口进入第二层循环得题目，一般都要在末尾加上这个符号
//如ip得题目要加上.或者:

func reverseWords(s string) string {
	//很聪明的办法,给原字符加一个空格
	s = " " + s
	ansString := ""
	for i := len(s) - 1; i >= 0; {
        //1.如果碰到了不是空格得字符（结尾有可能有很多空格）
		if s[i] != ' ' {
            var j int
            //2.找到第一个空格
			for j = i - 1; j >= 0; j-- {
				if s[j] == ' ' {
					ansString += s[j+1:i+1] + " "
					break
				}
			}
			i = j
		} else {
			i--
		}
	}

	if ansString != "" {
		return ansString[0 : len(ansString)-1]
	}
	return ansString
}

// @lc code=end
