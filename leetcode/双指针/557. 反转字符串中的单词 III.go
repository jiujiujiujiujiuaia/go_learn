package main

import (
	"bytes"
	"fmt"
)

func reverseWords(s string) string {
	bytesArray := []byte(s + " ")
	reverseString(bytesArray)
	var resultBytes []byte
	buf := bytes.NewBuffer(resultBytes)
	for i, j := len(bytesArray)-1, len(bytesArray); i >= 0; i-- {
		if bytesArray[i] == ' ' {
			buf.Write(bytesArray[i+1 : j])
			buf.WriteByte(' ')
			j = i
		}
	}
	return buf.String()[0 : buf.Len()-1]
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i <= j; {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}
}

func main() {
	str := "Let's take LeetCode contest"
	str = ""
	fmt.Println(reverseWords(str))
}
