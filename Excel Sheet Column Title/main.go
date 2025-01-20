package main


func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber-- // Giảm giá trị cột đi 1 để bắt đầu từ 0
		result = string('A' + columnNumber % 26) + result// Thêm ký tự vào đầu chuỗi kết quả
		columnNumber /= 26 // Chia cột cho 26 để xử lý cột tiếp theo
	}
	return result
}

func main(){
	columnNumber := 28
	println(convertToTitle(columnNumber)) // Kết quả mong muốn là "A"
}