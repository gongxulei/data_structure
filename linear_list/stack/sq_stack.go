/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/16
 * +----------------------------------------------------------------------
 * |Time: 3:09 下午
 * +----------------------------------------------------------------------
 */

package stack

import (
	"errors"
	"fmt"
)

// 顺序栈 Sequential Stack 基于切片实现

type SqStack struct {
	top    int
	base   int
	sizeOf int
	data   []int
}

func InitSqStack(length int) (sqStack *SqStack) {
	sqStack = new(SqStack)
	sqStack.sizeOf = length
	sqStack.data = make([]int, length)
	return
}

func (stack *SqStack) Push(num int) bool {
	if stack.top == stack.sizeOf {
		fmt.Println("error: 空间不足")
		return false
	}
	stack.data[stack.top] = num
	stack.top++
	return true
}

func (stack *SqStack) Pop() (res int, err error) {
	if stack.top == 0 {
		err = errors.New("空栈")
		return
	}
	stack.top--
	res = stack.data[stack.top]
	stack.data[stack.top] = 0
	return
}

func (stack *SqStack) IsEmpty() bool {
	if stack.top == 0 {
		return true
	}
	return false
}

func (stack *SqStack) Length() int {
	return stack.top
}
