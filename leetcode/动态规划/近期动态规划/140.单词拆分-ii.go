package main

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start

//TODO 
//1.思考一下深度搜索（递归+备忘录怎么做这个题，以及网友说的如何剪植优化）
//2.这个题的特例在下方，导致内存爆掉，如何优雅的处理（有人说自顶向下，研究下常用的自底和自顶区别）
func wordBreak(s string, wordDict []string) []string {
	wordDictMap := make(map[string]string)
	for _, word := range wordDict {
		wordDictMap[word] = word
	}

	judge := make([]bool, len(s)+1)
	dp := make([][]string, len(s)+1)
	judge[0] = true
	dp[0] = []string{""}
	

	//代码有点丑，先判断一遍，过掉那个特殊用例，然后再去保存
	for i := 1; i < len(s)+1;i++ {
		for j := 0; j < i; j++ {
			if judge[j] && wordDictMap[s[j:i]] != "" {
				judge[i] = true
				break
			}
		}
	}

	if !judge[len(s)]{
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

// func main(){
// 	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
//     ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]``
// }

// @lc code=end
