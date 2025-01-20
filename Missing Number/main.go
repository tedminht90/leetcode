package main

import "fmt"


func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}


func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}