package cn

import (
	"reflect"
	"testing"
)

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

func plusOne(digits []int) []int {
	l := len(digits) - 1
	digits[l] += 1
	for l >= 0 {
		if digits[l] >= 10 {
			digits[l] -= 10 // 减去10
			if l-1 < 0 {
				digits = append([]int{1}, digits...)
			} else {
				digits[l-1] += 1
			}
		}
		l--
	}
	return digits
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		data []int
		ex   []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{0}, []int{1}},
	}
	for k, test := range tests {
			if ret := plusOne(test.data); reflect.DeepEqual(ret, test.ex) {
				t.Log("测试第", k, "组数据成功", ret)
			} else {
				t.Log("测试第", k, "组数据失败", ret)
			}
		}
}
