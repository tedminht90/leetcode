package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	
	if sum%3 != 0 {
		return false
	}
	// Giá trị mục tiêu cho mỗi phần
	targetSum := sum / 3

	// Duyệt qua mảng và kiểm tra xem có thể chia thành 3 phần bằng nhau không
	currentSum := 0
    count := 0
	for _, num := range arr {
		currentSum += num
		if currentSum == targetSum {
			count++
			currentSum = 0
		}
	}
	// Cần tìm thấy ít nhất 3 phần có tổng bằng targetSum
    // Lưu ý: nếu có nhiều hơn 3 phần, cũng hợp lệ
	return count >= 3
}


func main(){
	arr := []int{0,2,1,-6,6,-7,9,1,2,0,1}
	println(canThreePartsEqualSum(arr))
}