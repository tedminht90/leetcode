package main


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    k := 1 // First element is always unique
    
    // Start from second element
    for i := 1; i < len(nums); i++ {
        // If current element is different from previous
        if nums[i] != nums[i-1] {
            nums[k] = nums[i] // Move unique element to position k
            k++
        }
    }
    return k
}

func main() {
	input := []int{0,0,1,1,1,2,2,3,3,4}
	println(removeDuplicates(input))
}

// Ví dụ chi tiết với nums = [1,1,2,2,3]:
// Bước 1: k = 1, i = 1
// nums = [1,1,2,2,3]
//           ↑
//           k i
// nums[1] == nums[0], bỏ qua
// Bước 2: k = 1, i = 2
// Copynums = [1,1,2,2,3]
//               ↑ ↑
//               k i
// nums[2] != nums[1]
// → nums[k] = nums[i]
// → [1,2,2,2,3]
// k++
// Bước 3: k = 2, i = 3
// nums = [1,2,2,2,3]
//             ↑ ↑
//             k i
// nums[3] == nums[2], bỏ qua
// Bước 4: k = 2, i = 4
// nums = [1,2,2,2,3]
//             ↑   ↑
//           k     i
// nums[4] != nums[3]
// → nums[k] = nums[i]
// → [1,2,3,2,3]
// k++
// Kết quả cuối cùng:

// nums = [1,2,3,2,3]
// k = 3 (độ dài mảng sau khi loại bỏ trùng lặp)
// Chỉ 3 phần tử đầu tiên [1,2,3] là quan trọng