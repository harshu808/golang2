package main

import (
	"fmt"
	"maps"
)

func main() {
	v := make(map[string]int)

	v["k1"] = 2
	v["k2"] = 4

	fmt.Println(v)

	n := map[string]int{"n1": 4, "n2": 3}

	fmt.Println(n)

	delete(v, "k1")
	fmt.Println(v)

	j1 := v["k2"]
	fmt.Println(j1)

	s1 := map[string]int{"l1": 2, "l2": 3}

	s2 := map[string]int{"l1": 2, "l2": 3}

	if maps.Equal(s1, s2) {
		fmt.Println("s1==s2")
	}

}
