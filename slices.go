package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string

	s = make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("s :", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s ", s)

	var c []string

	c = make([]string, len(s))
	copy(c, s)
	fmt.Println("c ", c)
	fmt.Println("s ", s)
	var j []string
	j = s[1:3]
	fmt.Println(j)

	var l []string
	l = s[:2]
	fmt.Println(l)

	var w1 []string

	w1 = []string{"s", "r", "h"}
	fmt.Println((w1))

	var w2 []string

	w2 = []string{"j", "l", "n"}
	fmt.Println(w2)

	if slices.Equal(w1, w2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	twod := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner := i + 1
		twod[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			fmt.Println("i+j = ", i+j)
		}
	}

}
