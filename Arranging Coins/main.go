package main

import "math"


func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(8*n+1))-1)/2)
}

func main(){
	println(arrangeCoins(5))
}
