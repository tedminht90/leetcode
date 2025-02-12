package main


func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i:] + s[:i] == goal {
			return true
		}
	}

	return false
}

func main(){
	s := "abcde"
	goal := "cdeab"

	println(rotateString(s, goal))
}