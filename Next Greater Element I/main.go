package main

import "fmt"


func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// Tạo map
	nextGreater := make(map[int]int)

	// Tìm next greater element cho tất cả các phần tử trong nums2
	for i := 0; i < len(nums2); i++ {
		nextGreater[nums2[i]] = -1 //Mặc định là -1
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				nextGreater[nums2[i]] = nums2[j]
				break
			}
		}
	}

	// Tạo mảng kết quả
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
}

func main(){
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}

	fmt.Println(nextGreaterElement(nums1, nums2))
}