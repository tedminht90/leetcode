package main


func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n % i == 0 {
			if check(s, i) {
				return true
			}
		}
	}
	return false
}
func check(s string, i int) bool {
	n := len(s)
	for j := i; j < n; j++ {
		if s[j] != s[j % i] {
			return false
		}
	}
	return true
}

func main(){
	s := "abab"
	println(repeatedSubstringPattern(s))
}