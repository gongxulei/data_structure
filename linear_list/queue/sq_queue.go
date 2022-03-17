/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/16
 * +----------------------------------------------------------------------
 * |Time: 5:56 下午
 * +----------------------------------------------------------------------
 */

package queue

import (
	"errors"
	"fmt"
)

// 顺序队列 Sequential Queue
// 循环队列 Circular Queue 顺序队列的优化版本，解决了空间浪费的问题

type SqQueue struct {
	// 队头(出队)
	front int
	// 队尾(入队)
	rear int
	data []int
}

func InitSqQueue(length int) (sqQueue *SqQueue) {
	if length <= 0 {
		panic("length invalid")
	}
	sqQueue = new(SqQueue)
	sqQueue.data = make([]int, length)
	return
}

func (sqQueue *SqQueue) Push(data int) bool {
	rear := (sqQueue.rear + 1) % len(sqQueue.data)
	if rear == sqQueue.front {
		fmt.Println("queue is full")
		return false
	}
	sqQueue.data[sqQueue.rear] = data
	sqQueue.rear = rear
	return false
}

func (sqQueue *SqQueue) Pop() (res int, err error) {
	if sqQueue.rear == sqQueue.front {
		// fmt.Println("queue is full")
		return 0, errors.New("queue is empty")
	}
	res = sqQueue.data[sqQueue.front]
	sqQueue.data[sqQueue.front] = 0
	sqQueue.front++
	return
}

func (sqQueue *SqQueue) Length() (length int) {
	length = (sqQueue.rear - sqQueue.front + len(sqQueue.data)) % len(sqQueue.data)
	return
}
