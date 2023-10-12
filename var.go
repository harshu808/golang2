package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for i := range nums {
		total += i

	}
	fmt.Println(total)
}

func main() {
	sum(2, 3, 45, 4)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
