package main


func findTheDifference(s, t string) byte {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if m[r] == 0 {
			return byte(r)
		}
		m[r]--
	}
	return 0
}

func main(){
	s := "abcd"
	t := "abcde"
	println(findTheDifference(s, t))
}