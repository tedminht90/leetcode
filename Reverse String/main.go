package main

import "fmt"

func reverseString(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}


func main(){
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reversed := reverseString(s)
	for _, char := range reversed{
		fmt.Printf("%c", char)
	}

}