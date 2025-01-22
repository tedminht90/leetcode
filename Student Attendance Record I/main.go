package main

func checkRecord(s string) bool {
	absent := 0
	late := 0
	for _, c := range s {
		if c == 'A' {
			absent++ 
			late = 0
		} else if c == 'L' {
			late++
		} else {
			late = 0
		}
		if absent > 1 || late > 2 {
			return false
		}
	}
	return true
}

func main() {
	s := "PPALLP"
	println(checkRecord(s))
}