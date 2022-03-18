/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/3/17
 * +----------------------------------------------------------------------
 * |Time: 7:26 下午
 * +----------------------------------------------------------------------
 */

package string

import (
	"testing"
)

func TestBFMatchString(t *testing.T) {
	strings := "aacbaacggjjkaacaacbaacbaacfffbaacbaacffdbaacb"
	s := "aacbaacb"
	index, counts, err := BFMatchString(strings, s)
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("匹配成功，位置：%v, 匹配次数：%v", index, counts)
	}
}

func TestKMPgetNext(t *testing.T) {
	s := "aacbaacb"
	for i := 0; i < len(s); i++ {
		next, maxlen := getNext(s, i)
		t.Logf("i=%v, next=%v, max_common_len=%v", i, next, maxlen)
	}
}

func TestKMPMatchString(t *testing.T) {
	strings := "aacbaacggjjkaacaacbaacbaacfffbaacbaacffdbaacb"
	s := "aacbaacb"
	index, counts, err := KMPMatchString(strings, s)
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("匹配成功，位置：%v, 匹配次数：%v", index, counts)
	}
}
