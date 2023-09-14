package AlgorithmsAndDataStructure

import "fmt"

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
}

type Stack struct {
	items []int
}

// Push
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) pop() int {
	l := len(s.items) - 1
	lastItem := s.items[l]
	s.items = s.items[:l]
	return lastItem
}
