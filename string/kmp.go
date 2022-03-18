/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/17
 * +----------------------------------------------------------------------
 * |Time: 4:50 下午
 * +----------------------------------------------------------------------
 */

package string

import (
	"errors"
)

// KMP匹配算法

// getNext 获取`模式字符串`的指针位置，最大公共长度
func getNext(s string, index int) (next int, maxCommonLen int) {
	if index <= 1 {
		return
	}
	SSlice := []rune(s)
	SSlice = SSlice[:index]
	j := 1
	for i := len(SSlice) - 1; i >= 0; i-- {
		if string(SSlice[:i]) == string(SSlice[j:]) {
			maxCommonLen = len(SSlice[:i])
			break
		}
		j++
	}
	if maxCommonLen != 0 {
		next = index - maxCommonLen - 1
	}
	return
}

func KMPMatchString(S string, T string) (index int, counts int, err error) {
	// string 转为切片
	SSlice := []rune(S)
	TSlice := []rune(T)
	if len(SSlice) < len(TSlice) {
		err = errors.New("模式串太长")
		return
	}
	j := index
	for i := 0; i < len(TSlice); i++ {
		counts++
		if j > (len(SSlice) - 1) {
			err = errors.New("目标串中未匹配到模式串")
			return
		}
		if TSlice[i] == SSlice[j] {
			j++
			continue
		}

		if i == 0 {
			j++
			index = j
			i = -1
			continue
		}

		// 匹配失败 初始化i位置
		next, maxCommonLen := getNext(T, i)
		if maxCommonLen == 0 {
			index = j
			i = -1
		} else {
			i = next - 1
			index += maxCommonLen + 1
		}
	}
	return
}
