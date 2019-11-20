package cn

import (
	"fmt"
	"testing"
)

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 你可以假设数组中无重复元素。
//
// 示例 1:
//
// 输入: [1,3,5,6], 5
//输出: 2
//
//
// 示例 2:
//
// 输入: [1,3,5,6], 2
//输出: 1
//
//
// 示例 3:
//
// 输入: [1,3,5,6], 7
//输出: 4
//
//
// 示例 4:
//
// 输入: [1,3,5,6], 0
//输出: 0
//
// Related Topics 数组 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	// 同理 双指针, 别瞎指挥解题
	if lens := len(nums); lens > 0 {
		left, right := 0, lens-1
		var mid int
		for left <= right {
			mid = (left + right) / 2 // 计算中间位置
			if nums[mid] > target {
				right = mid - 1 // 如果中间位置的值大于目标值就吧结束位置重置到中间位置前
			} else if nums[mid] < target {
				left = mid + 1 // 如果中间位置的值小于目标值就把开始位置设置到中间位置
			} else {
				return mid
			}
		}
		return mid
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInsertPosition(t *testing.T) {
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5) == 2)
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4) == 2)
}
