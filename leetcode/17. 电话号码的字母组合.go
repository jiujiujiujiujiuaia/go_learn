package main

import "fmt"

var maps = map[int][]string{2: {"a", "b", "c"}, 3: {"d", "e", "f"}, 4: {"g", "h", "i"}, 5: {"j", "k", "l"}, 6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"}, 8: {"t", "u", "v"}, 9: {"w", "x", "y", "z"}}

var res = make([]string, 0)

func letterCombinations(digits string) []string {
	if digits != "" {
		help(digits, "")
	}
	return res
}

func help(str string, com string) {
	if str == "" {
		res = append(res, com)
		return
	}

	for _, v := range maps[int(str[0])-48] {
		help(str[1:], com+v)
	}
}

func main() {
	fmt.Println(len(letterCombinations("323")))
}
