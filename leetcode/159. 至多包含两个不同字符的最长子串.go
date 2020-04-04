package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	substringMap := make(map[byte]int)
	var cnt int
	var startPoint = 0
	var maxSubStringLength = 0
	for i := 0; i < len(s); i++ {
		if _, ok := substringMap[s[i]]; !ok {
			cnt++
			substringMap[s[i]] = 1
		} else {
			substringMap[s[i]]++
		}

		if cnt == 3{
			if i - startPoint + 1 > maxSubStringLength{
				maxSubStringLength = i - startPoint + 1
			}

			for cnt == 3{
				if val,ok := substringMap[s[startPoint]];ok&&val-1==0{
					cnt--
				}
				startPoint++
			}
		}
	}
	return maxSubStringLength
}
