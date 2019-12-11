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
	//向右旋转 1 步: [7,1,2,3,4,5,6]
	//向右旋转 2 步: [6,7,1,2,3,4,5]
	//向右旋转 3 步: [5,6,7,1,2,3,4]
	var prevVal int
	var started bool
	for i, l := 0, len(nums); i < k; i++ {
		for j := 1; j <= l; j++ {
			// 记录要被覆盖的值
			if !started {
				prevVal = nums[j % l]
				nums[j%l] = nums[(j-1)%l]
				started = true
			} else {
				nums[j%l], prevVal = prevVal, nums[j%l]
			}
			fmt.Println(nums)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()


	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRotateArray(t *testing.T) {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(num, 3)
	fmt.Println(num)
}
