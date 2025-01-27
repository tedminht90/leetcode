package main

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	res := 0
	for k, v := range m {
		if m[k+1] != 0 {
			res = max(res, v + m[k+1])
		}
	}
	return res
}

func main(){
	nums := []int{1,3,2,2,5,2,3,7}
	println(findLHS(nums))
}