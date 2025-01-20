package main

import (
	"fmt"
	"math"
)


func constructRectangle(num int) []int {
	for i := int(math.Sqrt(float64(num))); i > 0; i-- {
		if num % i == 0 {
			return []int{num / i, i}
		}
	}
	return []int{num, 1}
}

func main(){
	area := 4
	fmt.Println(constructRectangle(area))
}