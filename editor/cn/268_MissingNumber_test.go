package cn

import (
	"fmt"
	"testing"
)

//给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。 
//
// 示例 1: 
//
// 输入: [3,0,1]
//输出: 2
// 
//
// 示例 2: 
//
// 输入: [9,6,4,2,3,5,7,0,1, 10]
//输出: 8
// 
//
// 说明: 
//你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现? 
// Related Topics 位运算 数组 数学

//leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	// 高斯求和
	//l, numsTotal := len(nums), 0
	//total := (1 + l) * l / 2 // 完整数列求和, 忽略0值, 但长度是n
	//for _, num := range nums {
	//	numsTotal += num	//缺失数字的数列求和
	//}
	//return total - numsTotal

	// 利用位运算特性  x ^ y ^ y = x
	num := len(nums)	// 基于n
	for idx := range nums {
		num ^= idx ^ nums[idx]	//缺失数字的数列求和
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1, 10}))
}
