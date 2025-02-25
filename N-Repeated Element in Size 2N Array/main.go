package main


func repeatedNTimes(nums []int) int {
    // Create a map to count occurrences
    count := make(map[int]int)
    
    // Tính n = độ dài mảng /2
    n := len(nums) / 2

	// Duyệt qua từng phần tử trong mảng
    for _, num := range nums {
        count[num]++
        // Early return when we find the element repeated n times
        if count[num] == n {
            return num
        }
    }
    
    return 0  // Should never reach here given problem constraints
}

func main(){
	nums := []int{1,2,3,3}
	println(repeatedNTimes(nums))
}