package main

import "fmt"


func isAnagram(s, t  string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

}