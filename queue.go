package queue

import "container/list"

type Queue struct{
    lock chan int
    list *list.List
}

func NewQueue() *Queue{
    lock := make(chan int, 1)
    list := list.New()
    return &Queue{lock, list}
}

func (q *Queue) Size() int{
    return q.list.Len()
}

func (q *Queue) Push(value interface{}){
    q.lock <- 1
    q.list.PushFront(value)
    <-q.lock
}

func (q *Queue) Pop() *list.Element{
    q.lock <- 1
    e := q.list.Back()
    q.list.Remove(e)
    <-q.lock
    return e
}