package main

import "fmt"

func main() {
	for i := 2; i < 4; i++ {
		fmt.Println(i)
	}
	for j := 4; j < 8; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
