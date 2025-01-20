package main


func canWinNim(n int) bool {
	return n % 4 != 0
}

func main() {
	n := 3
	println(canWinNim(n))	
}