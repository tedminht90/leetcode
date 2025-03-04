package main

func commonChars(words []string) []string {
	// Tạo mảng đếm ký tự xuất hiện trong từng từ
	charCount := make([][]int, len(words))
	for i, word := range words {
		charCount[i] = make([]int, 26)
		for _, char := range word {
			charCount[i][char-'a']++
		}
	}
	
	// Tạo mảng kết quả
	var res []string
	for i := 0; i < 26; i++ {
		minCount := 101
		for _, count := range charCount {
			if count[i] < minCount {
				minCount = count[i]
			}
		}
		for j := 0; j < minCount; j++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}


func main(){
	words := []string{"bella","label","roller"}
	println(commonChars(words))
}