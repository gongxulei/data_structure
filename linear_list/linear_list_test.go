/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/16
 * +----------------------------------------------------------------------
 * |Time: 3:16 下午
 * +----------------------------------------------------------------------
 */

package main

import (
	"data_structure/linear_list/linked_list"
	"data_structure/linear_list/queue"
	sq_stack "data_structure/linear_list/stack"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := linked_list.InitList()
	// fmt.Println(list)
	list.InsertNode(1, "test1")
	list.InsertNode(2, "test2")
	list.InsertNode(1, "test3")
	list.InsertNode(3, "test4")
	list.DeleteNode(2)
	fmt.Println(list.GetNode(2))
}

func TestSqStack(t *testing.T) {
	stack := sq_stack.InitSqStack(50)
	for i := 1; i <= 20; i++ {
		stack.Push(i)
	}
	t.Log("SqStack length:", stack.Length())
	for i := 0; i < 20; i++ {
		res, _ := stack.Pop()
		t.Log(res)
	}
}

func TestSqQueue(t *testing.T) {
	sqQueue := queue.InitSqQueue(10)
	for i := 1; i <= 10; i++ {
		sqQueue.Push(i)
	}
	// t.Log("SqQueue length:", sqQueue.Length())
	for i := 0; i < 10; i++ {
		res, _ := sqQueue.Pop()
		t.Log(res)
		sqQueue.Push(res)
	}
	t.Log(sqQueue)
}
