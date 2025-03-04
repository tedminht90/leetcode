package main

func findJudge(n int, trust [][]int) int{
	// Tạo mảng đếm số lượng người tin tưởng và số lượng người được tin tưởng
	trustCount := make([]int, n+1)
	trustedCount := make([]int, n+1)
	
	// Duyệt qua mảng trust, cập nhật thông tin
	for _, t := range trust{
		trustCount[t[0]]++
		trustedCount[t[1]]++
	}
	
	// Duyệt qua mảng đếm, tìm người được tin tưởng
	for i := 1; i <= n; i++{
		if trustCount[i] == 0 && trustedCount[i] == n-1{
			return i
		}
	}
	
	return -1
}

func main(){
	n := 2
	trust := [][]int{{1,2}}
	println(findJudge(n, trust))
}