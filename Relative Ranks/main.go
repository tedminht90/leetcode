package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	scoreMap := make(map[int]int)
	for i, v := range score {
		scoreMap[v] = i
	}
	
	sort.Sort(sort.Reverse(sort.IntSlice(score)))

	res := make([]string, len(score))
	for i, v := range score {
		switch i {
		case 0:
			res[scoreMap[v]] = "Gold Medal"
		case 1:
			res[scoreMap[v]] = "Silver Medal"
		case 2:
			res[scoreMap[v]] = "Bronze Medal"
		default:
			res[scoreMap[v]] = fmt.Sprintf("%d", i+1)
		}
	}
	return res
}




func main(){
	score := []int{5,4,3,2,1}
	fmt.Println(findRelativeRanks(score))
}