package main


func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for num % 2 == 0 {
		num /= 2
	}

	for num % 3 == 0 {
		num /= 3
	}

	for num % 5 == 0 {
		num /= 5
	}

	return num == 1
}

func main () {
	input := []int{6, 8, 14}

	for i := 0; i < len(input); i++ {
		println(isUgly(input[i]))
	}
}