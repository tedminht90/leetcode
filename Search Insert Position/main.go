package main



func searchInsert(nums []int, target int) int{
	for i  := 0; i< len(nums);i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	// nums := []int{1,3,5,6}
	// target := 5
	// nums := []int{1,3,5,6}
	// target := 2
	nums := []int{1,3,5,6}
	target := 7
	result := searchInsert(nums, target)
	println(result)
}