package main


func findDisappearedNumbers(nums []int) []int {
    // Đánh dấu số đã xuất hiện
    for _, num := range nums {
        index := abs(num) - 1
        nums[index] = -abs(nums[index])
    }
    
    // Thu thập các số chưa xuất hiện
    res := make([]int, 0)
    for i := range nums {
        if nums[i] > 0 {
            res = append(res, i + 1)
        }
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main(){
	nums := []int{1,11,10,7,3,6,9}
	println(findDisappearedNumbers(nums))
}