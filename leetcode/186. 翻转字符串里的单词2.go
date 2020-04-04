package main

func reverseWords(s []byte) {
	//1.先翻转整体
	helpReverseWords(s, 0, len(s)-1)
	//2.再翻转每一个单词
	for i := 0; i < len(s); {
		var j int
		for j = i + 1; j <= len(s); j++ {
			if j == len(s) || s[j] == ' ' {
				helpReverseWords(s, i, j-1)
				break
			}
		}
		//题目说了只有一个空格
		i = j + 1
	}
}

func helpReverseWords(s []byte, start, end int) {
	for end > start {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
