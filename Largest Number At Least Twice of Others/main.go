package main

import "fmt"


func dominantIndex(nums []int) int {
	max := 0
	secondMax := 0
	index := 0
	for i, num := range nums {
		if num > max {
			secondMax = max
			max = num
			index = i
		} else if num > secondMax {
			secondMax = num
		}
	}
	if max >= 2 * secondMax {
		return index
	}
	return -1
}

func main(){
	nums := []int{3, 4, 1, 8}
	fmt.Println(dominantIndex(nums))
}