package main


func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n % 3 == 0 {
		n /= 3
	}
	return n == 1
}

func main(){
	n := 27
	println(isPowerOfThree(n))
}