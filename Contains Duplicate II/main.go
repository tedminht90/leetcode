package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			if i - m[num] <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}

func main(){
	nums := []int{1,2,3,1,2,3}
	println(containsNearbyDuplicate(nums, 2))
}