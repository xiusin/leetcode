package cn

import (
	"fmt"
	"testing"
)

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
// 输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//
// 说明:
//
//
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。
//
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	// 数组想成一个环状, 移动指针就是将所有位置向后移动
	// 移动所有位置
	// 输入: [1,2,3,4,5,6,7] 和 k = 3
	//输出: [5,6,7,1,2,3,4]
	//解释:
	//正常不移动   : [1,2,3,4,5,6,7]
	//最终结果     : [5,6,7,1,2,3,4]
	fmt.Println(nums)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	l := len(nums)
	if l > 0 {
		k = k % l
		// 判断合适的移动方向
		var prevVal int
		willReplaceIdx := k % l
		prevVal = nums[willReplaceIdx]
		nums[willReplaceIdx] = nums[0]
		fmt.Println(nums)
		for i := k; i < l + k ; i++ {
			willReplaceIdx = ((i + 1) * k) % l
			fmt.Println("将要被覆盖的索引", willReplaceIdx)
			nums[willReplaceIdx], prevVal = prevVal, nums[willReplaceIdx]
			fmt.Println(nums)
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func TestRotateArray(t *testing.T) {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(num, 3)
	fmt.Println(num)
}
