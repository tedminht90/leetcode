package main


func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}

func main(){
	nums := []int{1,2,3,4}
	println(containsDuplicate(nums))
}