package main

import (
	"sort"
)


func largestSumAfterKNegations(nums []int, k int) int {
    // Bước 1: Sắp xếp mảng theo giá trị tuyệt đối giảm dần
    sortByAbsoluteValueDesc(nums)

    
    // Bước 2: Duyệt qua mảng và đổi dấu các số âm khi k > 0
    for i := 0; i < len(nums) && k > 0; i++ {
        if nums[i] < 0 {
            nums[i] = -nums[i]
            k--
        }
    }
    
    // Bước 3: Nếu k còn dư và là số lẻ, đổi dấu phần tử nhỏ nhất
    // (lúc này phần tử nhỏ nhất sẽ ở cuối mảng vì đã sắp xếp)
    if k%2 == 1 {
        nums[len(nums)-1] = -nums[len(nums)-1]
    }
    
    // Bước 4: Tính tổng mảng
    sum := 0
    for _, num := range nums {
        sum += num
    }
    
    return sum
}

// Hàm sắp xếp mảng theo giá trị tuyệt đối giảm dần
func sortByAbsoluteValueDesc(nums []int) {
    sort.Slice(nums, func(i, j int) bool {
        return abs(nums[i]) > abs(nums[j])
    })
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main(){
	nums := []int{4,2,3,-9}
	k := 1
	println(largestSumAfterKNegations(nums, k))
}