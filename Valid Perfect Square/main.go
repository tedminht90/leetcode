package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 0, num
	for l < r {
		mid := (l + r) / 2
		if mid * mid == num {
			return true
		} else if mid * mid < num {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}


func main(){
	num := 16
	fmt.Println(isPerfectSquare(num))
}