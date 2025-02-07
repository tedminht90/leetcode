package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividing(num int) bool {
	n := num
	for n > 0 {
		digit := n % 10
		if digit == 0 || num % digit != 0 {
			return false
		}
		n /= 10
	}
	return true
}


func main(){
	left := 1
	right := 22
	fmt.Println(selfDividingNumbers(left, right))
}