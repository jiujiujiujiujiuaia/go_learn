package main

import "math"

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (38.20%)
 * Likes:    177
 * Dislikes: 0
 * Total Accepted:    17.2K
 * Total Submissions: 44.7K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */

// @lc code=start
var min int

//解法一 纯回溯递归超时了，栈太深了
func ladderLength(beginWord string, endWord string, wordList []string) int {
	min = math.MaxInt32
	backTraceLadderLength(beginWord, endWord, wordList, 1)
	if min == math.MaxInt32 {
		return 0
	}
	return min
	// best := 0
	// for _,word := range wordList{
	// 	if judgeTwoWordIsOk(word, beginWord) {
	// 		backTraceLadderLength(word,endWord,)
	// 	}
	// }
}

func backTraceLadderLength(beginWord string, endWord string, wordList []string, cnt int) {
	if judgeTwoWordIsOk(beginWord, endWord) {
		for _, word := range wordList {
			if word == endWord {
				if min > cnt {
					min = cnt
				}
				return
			}
		}
	}

	for index, word := range wordList {
		if judgeTwoWordIsOk(word, beginWord) {
			wordListTemp := make([]string, len(wordList))
			copy(wordListTemp, wordList)
			backTraceLadderLength(word, endWord, append(wordListTemp[:index], wordListTemp[index+1:]...), cnt+1)
		}
	}

}

func judgeTwoWordIsOk(word1, word2 string) bool {
	cnt1 := 0
	for index, _ := range word1 {
		if word2[index] == word1[index] {
			cnt1++
		}
	}
	return (len(word1) - 1) == cnt1
}

// @lc code=end
