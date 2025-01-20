package main

import "fmt"


func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	res := 0
	odd := 0
	for _, v := range m {
		if v % 2 == 0 {
			res += v
		} else {
			res += v - 1
			odd = 1
		}
	}
	return res + odd
}

func main(){
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}