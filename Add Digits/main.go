package main


func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num % 9 == 0 {
		return 9
	}
	return num % 9
}


func main() {
	input := []int{38, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	println(addDigits(input[0]))
}