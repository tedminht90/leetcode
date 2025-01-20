package main

import "fmt"


func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
func main(){
	n := 11
	fmt.Println(hammingWeight(n))
}