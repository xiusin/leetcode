package cn

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
// 使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
//注意：
//
//答案中不可以包含重复的四元组。
//
//示例：
//
//给定数组 nums = [1, 0, -1, 0, -2, 2,  -]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]
//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	var arr [][]int
	l := len(nums)
	if l < 4 {
		return arr
	}
	// 先排序nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFourSum(t *testing.T) {
	// 提交时的错误时的测试用例
	tests := []struct {
		nums   []int
		target int
		except [][]int
	}{}
	fmt.Println(tests, fourSum([]int{1, 0, -1, 0, -2, 2, -3, 3}, 0))
	//for k, test := range tests {
	//	if ret := isValid(test.data); ret == test.ex {
	//		t.Log("测试第", k, "组数据成功")
	//	} else {
	//		t.Log("测试第", k, "组数据失败")
	//	}
	//}
}
