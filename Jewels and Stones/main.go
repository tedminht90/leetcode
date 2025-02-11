package main

import "fmt"


func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	for _, j := range J {
		m[j] = true
	}
	count := 0
	for _, s := range S {
		if m[s] {
			count++
		}
	}
	return count
}

func main(){
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}