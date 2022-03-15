/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/15
 * +----------------------------------------------------------------------
 * |Time: 4:21 下午
 * +----------------------------------------------------------------------
 */

package main

import (
	"fmt"
)


// 单向链表

func main() {
	list := InitList()
	// fmt.Println(list)
	list.InsertNode(1, "test1")
	list.InsertNode(2, "test2")
	list.InsertNode(1, "test3")
	list.InsertNode(3, "test4")
	list.DeleteNode(2)
	fmt.Println(list.GetNode(2))
}

type Node struct {
	data interface{} // 数据域
	next *Node       // 指针域
}

// LinkedList 链表头结点
type LinkedList struct {
	length   int // 储存链表的长度
	headNode *Node
}


// InitList 初始化链表
func InitList() *LinkedList {
	node := new(Node)
	L := new(LinkedList)
	L.headNode = node
	return L
}

// InsertNode 插入节点
func (linkedList *LinkedList) InsertNode(index int, v interface{}) {
	if index <= 0 {
		fmt.Println("err")
		return
	} else {
		node := &Node{
			data: v,
			next: nil,
		}
		// 头结点
		pre := linkedList.headNode
		if index == 1 {
			node.next = pre
			// 替换首元结点
			linkedList.headNode = node
		} else {
			for count := 1; count < index-1; count++ {
				pre = pre.next
			}
			node.next = pre.next
			pre.next = node
		}
		linkedList.length++
	}
}

// DeleteNode 删除节点
func (linkedList *LinkedList) DeleteNode(index int) bool {
	if index <= 0 {
		fmt.Println("err")
		return false
	}
	pre := linkedList.headNode
	if index == 1 {
		linkedList.headNode = pre.next
	} else {
		for i := 1; i < index-1; i++ {
			pre = pre.next
		}
		pre.next = pre.next.next
	}
	linkedList.length--
	return true
}

// GetNode 获取节点数据
func (linkedList *LinkedList) GetNode(index int) interface{} {
	if index <= 0 {
		fmt.Println("err")
		return nil
	} else {
		pre := linkedList.headNode
		for j := 1; j < index; j++ {
			if pre != nil {
				pre = pre.next
			}
		}
		return pre.data
	}
}
