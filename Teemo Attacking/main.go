package main

import "fmt"


func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	total := 0
	for i := 1; i < len(timeSeries); i++ {
		total += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return total + duration
}

func main() {
	timeSeries := []int{1, 4}
	duration := 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))
}