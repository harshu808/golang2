package main

import "fmt"

func main() {
	a := []int{2, 3, 4}
	sum := 0

	for _, b := range a {
		sum += b
	}
	fmt.Println(sum)

	for i, j := range a {
		if j == 3 {

			fmt.Println("3:", i)
		}

	}

	kv := map[string]string{"a": "apple", "b": "orange"}

	for l, o := range kv {
		fmt.Println(l, ":", o)
	}

	for m, s := range kv {
		fmt.Println("keys:", m, "for", s)
	}

}
