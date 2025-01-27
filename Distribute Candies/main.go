package main


func distributeCandies(candyType []int) int {
	m := make(map[int]bool)
	for _, v := range candyType {
		m[v] = true
	}
	return min(len(m), len(candyType)/2)
}

func main() {
	candyType := []int{1,1,2,2,3,3}
	println(distributeCandies(candyType))
}