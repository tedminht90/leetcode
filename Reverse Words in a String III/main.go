package main

import "strings"


func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = reverse(word)
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s := "Let's take LeetCode contest"
	println(reverseWords(s))
}