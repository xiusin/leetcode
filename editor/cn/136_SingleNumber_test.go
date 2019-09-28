package cn

import "testing"

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。 
//
// 说明： 
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？ 
//
// 示例 1: 
//
// 输入: [2,2,1]
//输出: 1
// 
//
// 示例 2: 
//
// 输入: [4,1,2,1,2]
//输出: 4 
// Related Topics 位运算 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	// o(n)来实现
	// 异或运算 特点, 两次异或同一个值原值不变
	// 位运算知识
	//一个数和 0 做 XOR 运算等于本身：a⊕0 = a
	//一个数和其本身做 XOR 运算等于 0：a⊕a = 0
	//XOR 运算满足交换律和结合律：a⊕b⊕a = (a⊕a)⊕b = 0⊕b = b
	a := 0 //   如果我们对 0 和二进制位做 XOR 运算，得到的仍然是这个二进制位
	for i := range nums {
		a ^= nums[i]
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSingleNumber(t *testing.T) {

}
