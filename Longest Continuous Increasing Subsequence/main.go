package main


func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count > max {
			max = count
		}
	}

	return max
}

func main(){
	nums := []int{1,3,5,4,7}

	println(findLengthOfLCIS(nums))
}