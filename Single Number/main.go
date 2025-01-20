package main


func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR 2 số với nhau
	}
	return result
}


func main(){
	nums := []int{4,1,2,1,2}
	println(singleNumber(nums))
}