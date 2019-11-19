package cn

import (
	"fmt"
	"testing"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	//双头指针, 保证交换以后的相对顺序
	// i需要查找0的数字
	// j需要查找不为0的数字
	var i, j, l = 0, 1, len(nums)
	for i < l && j < l {
		// 找到第一个不为零的数字
		if nums[j] == 0 {
			j++
			continue
		}
		// 找到第一个0
		if nums[i] != 0 {
			//移动i的位置
			i++
			// 如果i指针已经赶上了j指针, 那么后移j的位置
			if i == j {
				j++
			}
			continue
		}
		// 交换位置
		nums[i], nums[j] = nums[j], nums[i]
		// 两个指针都向后移继续查找
		i++
		j++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMoveZeroes(t *testing.T) {
	var nums = []int{0, 0, 1, 0, 0, 3, 12}
	moveZeroes(nums)
	fmt.Print(nums)

}
