package cn

import (
	"fmt"
	"testing"
)

//你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。 
//
// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。 
//
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。 
//
// 示例: 
//
// 给定 n = 5，并且 version = 4 是第一个错误的版本。
//
//调用 isBadVersion(3) -> false
//调用 isBadVersion(5) -> true
//调用 isBadVersion(4) -> true
//
//所以，4 是第一个错误的版本。  
// Related Topics 二分查找

func isBadVersion(n int) bool {
	return n >= 4
}

// todo 自己做了两遍不明白
func FirstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2 // 不用 (left + right) / 2 防止异常溢出数据类型
		if isBadVersion(mid) {       // 是错误版本 向前查找 正确 -> 错
			right = mid // 非正确版本,可能mid本身说明mid 可能是第一个错误版本, 所以包含进下次查找数据
		} else { // 如果是正确版本, 后移中间位置 + 1
			left = mid + 1
		}
	}
	return left
}

func TestFirstBadVersion(t *testing.T) {
	fmt.Println(FirstBadVersion(9))
}
