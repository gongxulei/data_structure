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

// 暴力匹配 brute_force

func BFMatchString(S string, T string) (index int, counts int, err error) {
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
		if j > len(SSlice)-1 {
			err = errors.New("目标串中未匹配到模式串")
			return
		}
		if TSlice[i] == SSlice[j] {
			j++
			continue
		}
		// 初始化i位置
		i = -1
		index++
		j = index
	}
	return
}
