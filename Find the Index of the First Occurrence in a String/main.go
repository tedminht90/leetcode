package main

import "fmt"


func strStr(haystack string, needle string) int{
	if len(needle)== 0 {
		return 0
	}
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		if haystack[i] == needle[0] {
			match := true
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}


func main(){
	haystack := "a"
	needle := "a"
	result := strStr(haystack, needle)
	fmt.Println(result)
}

// haystack = "hello", needle = "ll"

// i = 0: 
// - 'h' == 'l'? Không → tiếp tục

// i = 1:
// - 'e' == 'l'? Không → tiếp tục

// i = 2:
// - 'l' == 'l'? Đúng → kiểm tra tiếp
// - j = 0: haystack[2] == needle[0] → true
// - j = 1: haystack[3] == needle[1] → true
// - j = 2: kết thúc vòng lặp
// - match vẫn true → return 2