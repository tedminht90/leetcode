package main

func detectCapitalUse(word string) bool {
    count := 0
 	for i := 0; i < len(word); i++ {
        if isUpper(word[i]) {
            count++
        }
    }
    
    return count == 0 || // tất cả là chữ thường
           count == len(word) || // tất cả là chữ hoa  
           (count == 1 && isUpper(word[0])) // chỉ chữ đầu viết hoa
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}


func main(){
	println(detectCapitalUse("USA"))
	println(detectCapitalUse("FlaG"))
	println(detectCapitalUse("Flag"))
}