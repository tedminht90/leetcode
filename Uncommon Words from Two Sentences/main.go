package main

import (
	"fmt"
	"strings"
)


func uncommonFromSentences(s1, s2 string) []string{
	m := make(map[string]int)
	for _, v := range append(strings.Split(s1, " "), strings.Split(s2, " ")...){
		m[v]++
	}
	var res []string
	for k, v := range m{
		if v == 1{
			res = append(res, k)
		}
	}
	return res
}

func main(){
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2))
}