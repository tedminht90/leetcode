package main


func thirdMax(nums []int) int {
	max1, max2, max3 := -1 << 63, -1 << 63, -1 << 63
	for _, num := range nums {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num < max1 && num > max2 {
			max3 = max2
			max2 = num
		} else if num < max2 && num > max3 {
			max3 = num
		}
	}
	if max3 == -1 << 63 {
		return max1
	}
	return max3
}

func main(){
	nums := []int{3, 2, 1}
	println(thirdMax(nums))
}