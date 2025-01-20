package main

import (
	"fmt"
	"math"
)



func findThreeLargestNumbers(array []int) []int {
	if len(array) < 3 {
		return nil
	}
	first := math.MinInt32 // số lớn nhất
	second := math.MinInt32 // số lớn thứ 2
	third := math.MinInt32 // số lớn thứ 3
	for _, num := range array {
		if num > first { // TH1: nếu num lớn hơn số lớn nhất
			third = second // số lớn thứ 3 == số lớn thứ 2
			second = first // số lớn thứ 2 == số lớn nhất cũ
			first = num // số lớn nhất == num
		} else if num > second { // TH2: nếu num lớn hơn số lớn thứ 2
			third = second // Số lớn thứ 3 == số lớn thứ 2 cũ
			second = num // Số lớn thứ 2 == num
		} else if num > third { // TH3: nếu num lớn hơn số lớn thứ 3
			third = num // Số lớn thứ 3 == num
		}
	}
	return []int{third, second, first}
}

func main(){
	arr := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(findThreeLargestNumbers(arr))
}