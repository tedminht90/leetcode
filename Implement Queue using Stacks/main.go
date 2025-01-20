package main

import "fmt"


type MyQueue struct {
    queue []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.queue = append(this.queue, x)
}


func (this *MyQueue) Pop() int {
    first := this.queue[0]
	this.queue = this.queue[1:]
	return first
}


func (this *MyQueue) Peek() int {
    return this.queue[0]
}


func (this *MyQueue) Empty() bool {
    return len(this.queue) == 0
}

func main() {
	input := []string{"MyQueue", "push", "push", "peek", "pop", "empty"}

	obj := Constructor()

	result := make([]interface{}, 0)

	for i, op := range input {
		switch op {
		case "MyQueue":
			result = append(result, nil)
		case "push":
			if i == 1{
				obj.Push(1)
				result = append(result, nil)
			} else if i == 2 {
				obj.Push(2)
				result = append(result, nil)
			}
		case "peek":
			result = append(result, obj.Peek())
		case "pop":
			result = append(result, obj.Pop())
		case "empty":
			result = append(result, obj.Empty())
		}
	}
	fmt.Println(result)

}