package main

//test
func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i <= j; {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}
}
