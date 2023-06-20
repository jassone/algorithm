package main

import "fmt"

func huiwen(str string) bool{
	bytes := []rune(str)
	len := len(bytes)
	res := true

	for i:=0;i< len/2 ;i++ {
		if bytes[i] != bytes[len-i-1]{
			res = false
			break
		}
	}
	return res
}

func main() {
	s := "中abccba中"
	result := huiwen(s)
	if result {
		fmt.Printf("%s 是回文",s)
	} else {
		fmt.Printf("%s 不是回文",s)
	}
}
