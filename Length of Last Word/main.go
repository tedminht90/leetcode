package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	result := ""
	for i := len(s)-1; i>=0;i--{
		if s[i] != ' '{
			result = string(s[i]) + result
			length++
		} else if length >0 {
			fmt.Println(result)
			return length

		}
	}
	return length
}




func main(){
	//s := "Hello World"
	s := "luffy is still joyboy"
	println(lengthOfLastWord(s))
}