package main

func reverseStr(s string, k int) string {
	if k == 1 {
		return s
	}
	// chuyển string thành slice byte để xử lý
	bs := []byte(s)
	// xử lý từng đoạn ký tự
	for i := 0; i < len(bs); i += 2 * k {
		// tìm vi trí cuối cùng của chuỗi cần đảo
		end := min(i + k - 1, len(bs) - 1)

		// con trỏ j từ đầu, l từ cuối con trỏ chuyển vào giữa
		j := i
		l := end
		// Đảo cho đến khi 2 con trỏ gặp nhau
		for j < l {
			bs[j], bs[l] = bs[l], bs[j]
			j++
			l--
		}
	}
	return string(bs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main(){
	println(reverseStr("abcdefg", 2)) // Expected: "bacdfeg"
	println(reverseStr("abcd", 2)) // Expected: "bacd"
	println(reverseStr("abcdefg", 8)) // Expected: "gfedcba"
}