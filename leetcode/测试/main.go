package main

import "fmt"

func main() {

	a := [][]int{{1, 1}}
	fmt.Println("main:", a)
	edit(a)
	fmt.Println("main:", a)
}

func edit(a [][]int) {
	a[0][0] = 100
	fmt.Println("func:", a)
	edit2(a)
}

func edit2(a [][]int) {
	a[0][1] = 100
	fmt.Println("func:", a)
}
