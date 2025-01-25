package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums) // Sắp xếp bởi vì nó đảm bảo số nhỏ trong mỗi cặp là lớn nhất có thể
	sum := 0
	for i:=0; i<len(nums); i+=2 {
		sum += nums[i]
	}
	return sum
}


func main(){
	nums := []int{1,4,3,2}
	fmt.Println(arrayPairSum(nums))
}