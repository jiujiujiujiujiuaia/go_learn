package main

import "container/list"

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

//解法一 bfs 把单词接龙的链路抽象成寻找最短路径

//使用map记录路径会有一种情况，就是失败的路径把需要用的单词给占住了？
//不会，假设这个单词是必经点，如果搜索到了这个单词，必定可以走向最终的答案

//！！1.重点：使用广度优先最大的好处，因为是一层一层的遍历，因此只要是发现的第一条路径，那就是最短的路径
//map的目的是避免出现回环，也就是a到了b，然后b又到了a，会死循环
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//1.预处理
	transformations, routeMap := preHandle(wordList)
	queue := list.New()
	queue.PushBack(map[string]int{beginWord: 1})

	for queue.Len() > 0 {
		head := queue.Front()
		queue.Remove(head)
		value := head.Value.(map[string]int)

		//2.队列中存储的是map，也可以是结构体，存储的是步数和单词
		for key, val := range value {

			//3.依次遍历单词每一位，看单词能够变成几种相似的
			for i := 0; i < len(key); i++ {
				patterm := key[:i] + "*" + key[i+1:]
				if words, ok := transformations[patterm]; ok {
					for _, word := range words {
						if word == endWord {
							return val + 1
						}

						if !routeMap[word] {
							routeMap[word] = true
							queue.PushBack(map[string]int{word: val + 1})
						}
					}
				}
			}
		}
	}
	return 0
}

func preHandle(wordList []string) (map[string][]string, map[string]bool) {
	res := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			patterm := word[:i] + "*" + word[i+1:]
			res[patterm] = append(res[patterm], word)
		}
	}

	routeMap := make(map[string]bool)
	return res, routeMap
}

// @lc code=end
