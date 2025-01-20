package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}
	fmt.Println(m)
	res := []int{}
	for _, v := range nums2 {
		if m[v] {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

func main(){
	nums1 := []int{1,2,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersection(nums1, nums2))
}