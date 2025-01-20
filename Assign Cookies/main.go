package main

import "fmt"


func findContentChildren(g,s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort(g)
	sort(s)
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main(){
	g := []int{1,2,3}
	s := []int{1,1}
	fmt.Println(findContentChildren(g, s))
}