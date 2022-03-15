/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/4
 * +----------------------------------------------------------------------
 * |Time: 8:57 上午
 * +----------------------------------------------------------------------
 */

package main

// {10, 9, 98, 8, 6, 100, 5, 3, 101, 102, 103, 104}

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	var h = len(arr)/2
	// for h <= len(arr)/3 {
	// 	h = h*3 + 1
	// }
	for gap := h; gap > 0; gap = gap/2 {
		for i := 0; i < len(arr)-gap; i++ {
			for j := 0; j < len(arr)/gap; j++ {
				key := i + j*gap
				if key+gap < len(arr) && arr[key] > arr[key+gap] {
					temp := arr[key]
					arr[key] = arr[key+gap]
					arr[key+gap] = temp
				}
			}

		}
	}

	return arr
}
