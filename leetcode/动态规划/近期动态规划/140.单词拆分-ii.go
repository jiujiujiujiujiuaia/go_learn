package main

import "fmt"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start

// TODO
//1.这个题的特例在下方，导致内存爆掉，如何优雅的处理（有人说自顶向下，研究下常用的自底和自顶区别）
func wordBreak1(s string, wordDict []string) []string {
	wordDictMap := make(map[string]string)
	for _, word := range wordDict {
		wordDictMap[word] = word
	}

	judge := make([]bool, len(s)+1)
	dp := make([][]string, len(s)+1)
	judge[0] = true
	dp[0] = []string{""}

	//代码有点丑，先判断一遍，过掉那个特殊用例，然后再去保存
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if judge[j] && wordDictMap[s[j:i]] != "" {
				judge[i] = true
				break
			}
		}
	}

	if !judge[len(s)] {
		return []string{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) != 0 && wordDictMap[s[j:i]] != "" {
				for _, word := range dp[j] {
					var newWord string
					if word == "" {
						newWord += wordDictMap[s[j:i]]
					} else {
						newWord = word + " " + wordDictMap[s[j:i]]
					}
					dp[i] = append(dp[i], newWord)
				}
			}
		}
	}

	return dp[len(s)]
}

//解法二：递归的解法
//递归的解法好处在于会先确保问题能够解决，也就是不会出现无法解决的例子浪费空间
func wordBreak(s string, wordDict []string) []string {
	//1.放在map中便于快速查找
	wordDictMap := make(map[string]string)
	for _, word := range wordDict {
		wordDictMap[word] = word
	}

	dp := make([][]string, len(s)+1)
	helpWordBreak(s, wordDictMap, 0, dp)
	return dp[0]
}

func helpWordBreak(s string, wordDictMap map[string]string, lastWordEnd int, dp [][]string) []string {

	if len(dp[lastWordEnd]) != 0 {
		return dp[lastWordEnd]
	}

	res := make([]string, 0)

	if lastWordEnd == len(s) {
		res = append(res, "")
		return res
	}

	for end := lastWordEnd + 1; end <= len(s); end++ {
		if wordDictMap[s[lastWordEnd:end]] != "" {
			//如果后面的字符无法拆分，则为空。则无法解决的问题不会浪费空间。
			wordList := helpWordBreak(s, wordDictMap, end, dp)
			for _, word := range wordList {
				if word == "" {
					res = append(res, s[lastWordEnd:end])
				} else {
					res = append(res, s[lastWordEnd:end]+" "+word)
				}
			}
		}
	}

	//递归的dp定义和迭代的dp定义相反
	dp[lastWordEnd] = append(dp[lastWordEnd], res...)
	fmt.Println(lastWordEnd, dp[lastWordEnd])
	return res
}

// func main(){
// 	"aaaaaaba"
//   ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]``
// }

// @lc code=end
