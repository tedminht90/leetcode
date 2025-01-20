package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// Lấy chuỗi đầu tiên làm tiền tố khởi đầu để so sánh
	prefix := strs[0]
	// Vòng lặp bên ngoài: duyệt qua từng chuỗi trong mảng, bắt đầu từ chuỗi thứ 2 (i = 1)
	for i := 1; i < len(strs); i++ {

		// Vòng lặp bên trong: duyệt đồng thời qua các ký tự của:
		// - prefix: chuỗi tiền tố hiện tại
		// - chuỗi đang xét strs[i]
		// Điều kiện j < len(prefix) && j < len(strs[i]) để tránh truy cập vượt quá độ dài của chuỗi
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {

			// So sánh ký tự tại vị trí j của prefix và chuỗi hiện tại
			// Nếu khác nhau thì cắt prefix từ vị trí j trở về sau
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		// đoạn code này xử lý thêm một trường hợp đặc biệt: khi prefix dài hơn chuỗi đang xét.
		for len(prefix) > len(strs[i]) {
            prefix = prefix[:len(prefix)-1]
        }
	}

	return prefix
}


func main(){
	var input = []string{"flower","flow","flight"}
	//var input = []string{"dog","racecar","car"}
	//var input = []string{"ab","a"}
	fmt.Println(longestCommonPrefix(input))

}

// Phiên bản mới:
// - prefix ban đầu = "flower"
// - So sánh với "flo":
//   + f giống nhau
//   + l giống nhau
//   + o giống nhau
//   + Vòng lặp dừng vì j đạt đến cuối "flo"
//   + Thêm vòng lặp mới cắt prefix về độ dài bằng "flo"
//   + prefix = "flo" (chính xác)