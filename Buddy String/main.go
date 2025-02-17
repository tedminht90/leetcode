package main

import "fmt"


func buddyStrings(s,goals string)bool {
	if len(s) != len(goals){
		return false
	}
	if s == goals{
		m := make(map[rune]bool)
		for _, c := range s{
			if _, ok := m[c]; ok{
				return true
			}
			m[c] = true
		}
		return false
	}
	var diff []int
	for i := 0; i < len(s); i++{
		if s[i] != goals[i]{
			diff = append(diff, i)
		}
	}
	return len(diff) == 2 && s[diff[0]] == goals[diff[1]] && s[diff[1]] == goals[diff[0]]
}

func main(){
	A := "aa"
	B := "aa"
	fmt.Println(buddyStrings(A, B))
}