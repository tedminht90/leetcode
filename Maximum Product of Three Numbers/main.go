package main

import "fmt"

func maximumProduct(nums []int) int {
	sorts := sort(nums)
	n := len(nums)
	
	return max(sorts[0]*sorts[1]*sorts[n-1], sorts[n-1]*sorts[n-2]*sorts[n-3])
	

}

func sort(num []int) []int{
	for i := 0; i < len(num); i++{
		for j := i; j < len(num); j++{
			if num[i] > num[j]{
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	return num
}


func main(){
	nums := []int{1,2,3,4}
	fmt.Println(maximumProduct(nums))

}