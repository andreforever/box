package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arrNum []int
	tam    int
}

func (q *Queue) Enqueue(num int) {
	if q.GetLength() == q.tam {
		fmt.Println("Overflow")
		return
	}
	q.arrNum = append(q.arrNum, num)
}

func (q *Queue) Dequeue() int {
	if q.Vacio() {
		fmt.Println("UnderFlow")
		return 0
	}
	elemento := q.arrNum[0]
	if q.GetLength() == 1 {
		q.arrNum = nil
		return elemento
	}
	q.arrNum = q.arrNum[1:]
	return elemento // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.arrNum)
}

func (q *Queue) Vacio() bool {
	return len(q.arrNum) == 0
}

func (q *Queue) Peek() (int, error) {
	if q.Vacio() {
		return 0, errors.New("empty queue")
	}
	return q.arrNum[0], nil
}

func main() {
	queue := Queue{tam: 3}
	fmt.Println(queue.arrNum)
	queue.Enqueue(11)
	fmt.Println(queue.arrNum)
	queue.Enqueue(22)
	fmt.Println(queue.arrNum)
	queue.Enqueue(33)
	fmt.Println(queue.arrNum)
	queue.Enqueue(55)
	fmt.Println(queue.arrNum)
	elemento := queue.Dequeue()
	fmt.Println(elemento)
	fmt.Println(queue.arrNum)
	queue.Enqueue(99)
	fmt.Println(queue.arrNum)
	elemento = queue.Dequeue()
	fmt.Println(elemento)
	fmt.Println(queue.arrNum)

}
