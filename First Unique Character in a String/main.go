package main

import "fmt"


func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	fmt.Println(m)
	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}
	return -1
}

func main(){
	s := "aab"
	println(firstUniqChar(s))
}