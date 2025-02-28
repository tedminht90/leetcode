package main

import (
	"fmt"
	"sort"
)


func largestPerimeter(A []int) int {
	// sort the array in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	for i := 0; i < len(A) - 2; i++{
		if A[i] < A[i + 1] + A[i + 2]{
			return A[i] + A[i + 1] + A[i + 2]
		}
	}
	return 0
}




func main(){
	nums := []int{2, 1, 2}
	fmt.Println(largestPerimeter(nums))
}