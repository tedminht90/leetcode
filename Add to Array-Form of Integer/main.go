package main

import "fmt"


func addToArrayForm(nums []int, k int)[]int{
	var res []int
	carry := 0
	for i := len(nums) - 1; i >= 0 || k > 0; i--{
		if i >= 0{
			carry += nums[i]
		}
		if k > 0{
			carry += k % 10
			k /= 10
		}
		res = append(res, carry % 10)
		carry /= 10
	}
	if carry > 0{
		res = append(res, carry)
	}
	reverse(res)
	return res
}

func reverse(nums []int){
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main(){
	nums := []int{1,2,0,0}
	k := 34
	fmt.Println(addToArrayForm(nums, k))
}