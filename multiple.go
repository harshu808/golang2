package main

import "fmt"

func vals() (int, int) {
	return 4, 5
}

func main() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)
}
