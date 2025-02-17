package main


func backspaceCompare(S string, T string) bool {
	return build(S) == build(T)
}

func build(S string) string{
	var stack []rune
	for _, c := range S{
		if c != '#'{
			stack = append(stack, c)
		}else if len(stack) > 0{
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func main(){
	S := "ab#c"
	T := "ad#c"
	println(backspaceCompare(S, T))
}