package Queue

import "fmt"

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(111)
	myQueue.enqueue(222)
	myQueue.enqueue(333)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}

type Queue struct {
	items []int
}

func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) dequeue() int {
	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem
}
