package main

import "fmt"

func main() {
	i := 2
	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if i < 0 {
		fmt.Println("negative")
	} else if i > 0 {
		fmt.Println("positive")

	} else {
		fmt.Println("any number")
	}
}
