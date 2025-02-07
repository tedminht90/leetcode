package main

import "fmt"



func pivotIndex(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    if len(nums) == 1 {
        return 0
    }
    
    sum := 0
    // Tính sum trong quá trình nhập mảng
    for _, num := range nums {
        sum += num
    }
    
    leftSum := 0
    for i, num := range nums {
        rightSum := sum - leftSum - num
        if leftSum == rightSum {
            return i
        }
        leftSum += num
    }
    return -1
}

func main(){
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))

}