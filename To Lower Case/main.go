package main

import "fmt"



func toLowerCase(str string) string {
	var result string
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func main(){

	strings := []string{"Hello", "here", "LOVELY"}

	for _, str := range strings {
		fmt.Println(toLowerCase(str))
	}
}