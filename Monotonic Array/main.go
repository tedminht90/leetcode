package main


func isMonotonic(nums []int) bool {
    state := 0  // 0: chưa xác định, 1: tăng, -1: giảm
    
    for i := 0; i < len(nums)-1; i++ {  // Chạy từ 0 đến n-2
        if nums[i] > nums[i+1] {  // So sánh với phần tử kế tiếp
            if state == 1 {        // Nếu đã xác định là tăng
                return false       // Mà gặp cặp giảm -> không đơn điệu
            }
            state = -1            // Đánh dấu là giảm
        } else if nums[i] < nums[i+1] {
            if state == -1 {       // Nếu đã xác định là giảm
                return false       // Mà gặp cặp tăng -> không đơn điệu
            }
            state = 1             // Đánh dấu là tăng
        }
    }
    
    return true  // Nếu không phát hiện mâu thuẫn -> đơn điệu
}

func main(){
	nums := []int{1, 2, 2, 2}
	println(isMonotonic(nums))
}