/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/9
 * +----------------------------------------------------------------------
 * |Time: 10:19 下午
 * +----------------------------------------------------------------------
 */

package main


// 堆排序
// 堆：完全二叉树(结构为从上到下，从左到右) + 父节点 > 子节点



func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func heap(tree []int, n, i int) {
	if i >= n {
		return
	}
	var (
		// parent = (i - 1) / 2
		c1  = 2*i + 1
		c2  = 2*i + 2
		max = i
	)
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}

	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		swap(tree, max, i)
		heap(tree, n, max)
	}
}

func buildHeap(tree []int, n int) {
	var (
		max    = n - 1
		parent = (max - 1) / 2
	)
	for i := parent; i >= 0; i-- {
		heap(tree, n, i)
	}
}

func HeapSort(tree []int) {
	buildHeap(tree, len(tree))
	for n := len(tree) - 1; n > 0; n-- {
		// 构建堆
		swap(tree, n, 0)
		heap(tree, n, 0)
	}
}
