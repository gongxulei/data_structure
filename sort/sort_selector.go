/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/3
 * +----------------------------------------------------------------------
 * |Time: 10:29 上午
 * +----------------------------------------------------------------------
 */

package main

import (
	"fmt"
)

func main() {

	var sortSlice = []int{3, 1, 6, 2, 10, 20, 13,100,70,30,56}
	// res := selectSort(sortSlice)
	// res := BubblingSort(sortSlice)
	// res := InsertSorted(sortSlice)
	// res := ShellSort(sortSlice)
	// res := MergeSort(sortSlice)
	HeapSort(sortSlice)
	fmt.Println(sortSlice)
	// fmt.Println(t())
}

// selectMin
func selectMin(slice []int, i int) (int, int) {
	var minPos = i
	var maxPos = len(slice) - i - 1
	fmt.Println(maxPos)
	for j := i; j < len(slice); j++ {
		if slice[j] < slice[minPos] {
			minPos = j
		}
		if j < len(slice)-i-1 && slice[j] > slice[maxPos] {
			maxPos = j
		}
	}
	return minPos, maxPos
}

// selectSort 选择排序：每次循环取出剩余数字中的最大值或最小值
// 时间复杂度 n + (n-1) + (n-2) + .... + 2 + 1 + 0 = n²/2  = O(n²)
// 空间复杂度：没有使用其他数组 O(1)
func selectSort(slice []int) []int {
	for i := 0; i < (len(slice))/2; i++ {
		minKey, maxKey := selectMin(slice, i)
		tmp := slice[i]
		slice[i] = slice[minKey]
		slice[minKey] = tmp

		tmpKey := len(slice) - 1 - i
		if maxKey == i {
			maxKey = minKey
		}
		// 高位交换
		tmp = slice[tmpKey]
		slice[tmpKey] = slice[maxKey]
		slice[maxKey] = tmp
	}
	return slice
}
