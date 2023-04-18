package main

import "fmt"

type IQueue interface {
	Enqueue(value int)
	Front()
	Unqueue()
}

type Queue struct {
	values []int
}

func (q *Queue) Enqueue(value int) {
	q.values = append(q.values, value)
}

func (q *Queue) Unqueue() int {
	if len(q.values) == 0 {
		return -1
	}
	primeiroValor := q.values[0]
	q.values = (q.values)[1:len(q.values)]
	return primeiroValor
}

func (q *Queue) Front() int {
	return q.values[0]
}

func main() {
	queue := Queue{}

	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)

	fmt.Println(queue.values)

	fmt.Println(queue.Front())

	queue.Unqueue()

	fmt.Println(queue.Front())

}
