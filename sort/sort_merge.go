/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/8
 * +----------------------------------------------------------------------
 * |Time: 9:49 上午
 * +----------------------------------------------------------------------
 */

package main

import (
	"fmt"
)


// {3, 1, 6, 2, 10, 20, 13}

// MergeSort 归并排序的入口
func MergeSort(arr []int) []int {
	// 分配一个辅助的数组
	var tempArr = make([]int, len(arr))
	sort(arr, tempArr, 0, len(arr)-1)
	return arr
}

func merge(arr, tempArr []int, left, mid, right int) {
	// fmt.Println(arr, mid, left, right)
	// 标记左半边未排序数字
	var lpos = left
	// 标记右半边未排序元素
	var rpos = mid + 1
	// 临时数组元素下标
	var pos = left
	// 合并
	for lpos <= mid && rpos <= right {
		if arr[lpos] < arr[rpos] { // 左半区第一个剩余元素更小
			tempArr[pos] = arr[lpos]
			pos++
			lpos++
		} else { // 右半区剩余的元素更小
			tempArr[pos] = arr[rpos]
			pos++
			rpos++
		}
	}

	for lpos <= mid {
		tempArr[pos] = arr[lpos]
		pos++
		lpos++
	}

	for rpos <= right {
		tempArr[pos] = arr[rpos]
		pos++
		rpos++
	}

	for left <= right {
		arr[left] = tempArr[left]
		left++
	}

	fmt.Println(tempArr, mid, left, right)
}

func sort(arr, tempArr []int, left, right int) {
	// 如果元素个数为1，那么就不需要继续划分
	if left == right {
		return
	}
	// 只有一个元素的区域，本身就是有序的，只需要被归并即可
	if left < right {
		// 中间点
		mid := (left + right) / 2
		// 递归划分左半区
		sort(arr, tempArr, left, mid)
		// 递归划分右半区
		sort(arr, tempArr, mid+1, right)
		// 合并已经排序的部分
		// fmt.Println(arr, left, mid, right)
		merge(arr, tempArr, left, mid, right)
	}
	return
}
