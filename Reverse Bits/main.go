package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// res <<= 1: Dịch trái res 1 bit để tạo chỗ cho bit mới
		// num & 1: Lấy bit cuối cùng (LSB) của số đầu vào bằng phép AND với 1
		// res |= num & 1: Gán bit vừa lấy được vào vị trí cuối của res bằng phép OR
		// num >>= 1: Dịch phải số đầu vào 1 bit để xét bit tiếp theo
		res <<= 1
		res |= num & 1
		num >>= 1
	}
	return res
}

func main() {
	input := uint32(0x94967295)
	output := reverseBits(input)

	fmt.Printf("Input (hex): 0x%x\n", input)
    fmt.Printf("Input (binary): %032b\n", input)
    fmt.Printf("Output (hex): 0x%x\n", output)
    fmt.Printf("Output (binary): %032b\n", output)
}