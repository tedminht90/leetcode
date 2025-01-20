package main

import (
	"fmt"
	"strings"
)


func wordPattern(pattern string, s string) bool {
    // Tách chuỗi s thành mảng các từ
    words := strings.Fields(s)
    
    // Kiểm tra độ dài pattern và số từ có bằng nhau không
    if len(pattern) != len(words) {
        return false
    }
    
    // Tạo 2 map để lưu ánh xạ
    patternToWord := make(map[byte]string)    // Lưu ký tự -> từ
    wordToPattern := make(map[string]byte)    // Lưu từ -> ký tự
    
    // Duyệt qua từng vị trí
    for i := 0; i < len(pattern); i++ {
        char := pattern[i]     // Ký tự trong pattern
        word := words[i]       // Từ tương ứng
        
        // Kiểm tra ánh xạ pattern -> word
        if mappedWord, exists := patternToWord[char]; exists {
            if mappedWord != word {
                return false
            }
        } else {
            patternToWord[char] = word
        }
        
        // Kiểm tra ánh xạ word -> pattern
        if mappedChar, exists := wordToPattern[word]; exists {
            if mappedChar != char {
                return false
            }
        } else {
            wordToPattern[word] = char
        }
    }
    
    return true
}

func main(){
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}