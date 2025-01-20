package main

import "fmt"

func removeElement(nums []int, val int) int{
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3,2,2,3}
	val := 3
	k := removeElement(nums, val)
	fmt.Println(k)
	fmt.Println(nums[:k])
}


// Input: nums = [3,2,2,3], val = 3
// Quá trình xử lý:
// - k = 0, num = 3: bỏ qua vì num = val
// - k = 0, num = 2: nums[0] = 2, k = 1
// - k = 1, num = 2: nums[1] = 2, k = 2
// - k = 2, num = 3: bỏ qua vì num = val

// Kết quả:
// - nums = [2,2,2,3] (chỉ quan tâm 2 phần tử đầu)
// - Trả về k = 2 (số phần tử còn lại)