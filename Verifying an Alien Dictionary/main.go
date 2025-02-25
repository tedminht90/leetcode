package main

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}
	
	for i := 0; i < len(words)-1; i++ {
		if !isSorted(words[i], words[i+1], orderMap) {
			return false
		}
	}
	
	return true
}

func isSorted(word1, word2 string, orderMap map[byte]int) bool {
	for i := 0; i < len(word1) && i < len(word2); i++ {
		if orderMap[word1[i]] < orderMap[word2[i]] {
			return true
		} else if orderMap[word1[i]] > orderMap[word2[i]] {
			return false
		}
	}
	
	return len(word1) <= len(word2)
}


func main(){
	words := []string{"hello","leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	println(isAlienSorted(words, order))
}