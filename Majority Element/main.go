package main


func majorityElement(nums []int) int {
	countMap := make(map[int]int)

	// Đếm số lần xuất hiện của mỗi số
	for _, num := range nums {
		countMap[num]++
	}
	// Tìm số xuất hiện nhiều nhất
	maxCount := 0
	result := 0
	for num, count := range countMap {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	return result
}

func main(){
	nums := []int{3,3,8,6,7,3}
	println(majorityElement(nums))
}