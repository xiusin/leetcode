package cn

import (
	"testing"
)

//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//
// 
//
// 在杨辉三角中，每个数是它左上方和右上方的数的和。 
//
// 示例: 
//
// 输入: 3
//输出: [1,3,3,1]
// 
//
// 进阶： 
//
// 你可以优化你的算法到 O(k) 空间复杂度吗？ 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2 MB,击败了92.68% 的Go用户
// 更牛逼的解法: https://leetcode-cn.com/problems/pascals-triangle-ii/solution/c-brief-solution-by-easier/
func getRow(rowIndex int) []int {
	var ret []int
	var p int
	ret = []int{1}
	for i := 1; i <= rowIndex; i++ {
		p = 0
		for j := 0; j < i-1; j++ {
			if p == 0 {
				p = ret[j+1]
				ret[j+1] = p + ret[j]
			} else {
				p, ret[j+1] = ret[j+1], p+ret[j+1]
			}
		}
		ret = append(ret, 1)
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPascalsTriangleIi(t *testing.T) {
	t.Log(getRow(0))
}
