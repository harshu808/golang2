package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus1(a, b, c int) int {
	return a + b + c
}

func main() {
	f1 := plus(1, 2)

	fmt.Println(f1)

	f2 := plus1(1, 2, 3)

	fmt.Println(f2)
}
