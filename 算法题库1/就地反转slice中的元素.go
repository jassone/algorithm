package main

import "fmt"

func revert(s []int) []int {
	length := len(s)
	for i := 0; i <= length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}

	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(revert(s))
}
