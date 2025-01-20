package main

import (
	"fmt"
	"strings"
)


func countSegments(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	words := strings.Split(s, " ")
	count := 0
	for _, word := range words {
		if word != "" {
			count++
		}
	}
	return count
}

func main(){
	s := "Hello, my name is John"
	fmt.Println(countSegments(s))
}