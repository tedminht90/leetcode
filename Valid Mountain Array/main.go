package main


func validMountainArray(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }
    
    left, right := 0, n-1
    
    // Di chuyển left sang phải nếu tăng
    for left+1 < n && nums[left] < nums[left+1] {
        left++
    }
    
    // Di chuyển right sang trái nếu giảm
    for right > 0 && nums[right-1] > nums[right] {
        right--
    }
    
    // Kiểm tra đỉnh trùng nhau và không ở biên
    return left == right && left != 0 && right != n-1 && left != n-1
}



func main() {
	nums := []int{0,2,3, 4,5, 2, 1, 0}
	println(validMountainArray(nums))
}