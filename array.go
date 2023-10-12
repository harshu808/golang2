package main

import "fmt"

func main() {
	var a [5]int

	b := [6]int{1, 2, 3, 4, 5, 6}

	var twod [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twod[i][j] = i + j

		}
	}
	fmt.Println(twod)
	fmt.Println(a)
	fmt.Println(b)
}
