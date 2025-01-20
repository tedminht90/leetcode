package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) -1; i >=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func main(){
	digits := []int{9}
	fmt.Println(plusOne(digits))
}