package node

import (
	"go-toolbox/queue/common"
	"sync"
)

var administrator = make(chan common.PutBody, 10)

func (q *Queue) PutChan(body common.PutBody) {
	administrator <- body
}
func (q *Queue) PutQueueChan() {
	for {
		select {
		case data := <-administrator:
			if data.Exigency {
				q.putHead(data.Data)
			} else {
				q.put(data.Data)
			}
			break
		default:
			break
		}
	}
}

// 代表每一个节点
type node struct {
	data interface{}
	next *node
}

type Queue struct {
	// 头节点
	head *node

	// 队尾节点
	rear *node

	size int

	sync.Mutex
}

func NewQueue() *Queue {
	q := new(Queue)
	q.head = nil
	q.rear = nil
	q.size = 0
	return q
}

// Put 尾插法
func (q *Queue) put(element interface{}) {
	n := new(node)
	n.data = element
	q.Lock()
	defer q.Unlock()

	if q.rear == nil {
		q.head = n
		q.rear = n
	} else {
		q.rear.next = n
		q.rear = n
	}
	q.size++
}

// PutHead 头插法，在队列头部插入一个元素
func (q *Queue) putHead(element interface{}) {
	n := new(node)
	n.data = element
	q.Lock()
	defer q.Unlock()
	if q.head == nil {
		q.head = n
		q.rear = n
	} else {
		n.next = q.head
		q.head = n
	}
	q.size++
}

// Get 获取并删除队列头部的元素
func (q *Queue) get() interface{} {
	if q.head == nil {
		return nil
	}
	n := q.head
	q.Lock()
	defer q.Unlock()
	// 代表队列中仅一个元素
	if n.next == nil {
		q.head = nil
		q.rear = nil

	} else {
		q.head = n.next
	}
	q.size--
	return n.data
}
