package main

import "fmt"

func findErrorNums(nums []int) []int {
    m := make(map[int]int)
    
    // Đếm số lần xuất hiện của mỗi số
    for _, num := range nums {
        m[num]++
    }
    
    duplicate, missing := 0, 0
    // Kiểm tra từ 1 đến n
    for i := 1; i <= len(nums); i++ {
        if m[i] == 2 {  // Số xuất hiện 2 lần
            duplicate = i
        }
        if m[i] == 0 {  // Số không xuất hiện
            missing = i
        }
    }
    
    return []int{duplicate, missing}
}


func main(){
	nums := []int{1,2,2,4}
	fmt.Println(findErrorNums(nums))
}