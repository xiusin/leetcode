package cn

import (
	"fmt"
	"testing"
)

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
// 
//
// 在杨辉三角中，每个数是它左上方和右上方的数的和。 
//
// 示例: 
//
// 输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//] 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.3 MB,击败了68.85% 的Go用户
func generate(numRows int) [][]int {
	var ret [][]int
	if numRows > 0 {
		ret = append(ret, []int{1})
		for i := 1; i < numRows; i++ {
			arr := []int{1}
			// 迭代上一个数组的数据
			for j := 0; j < i-1; j++ {
				arr = append(arr, ret[i-1][j]+ret[i-1][j+1])
			}
			// 追加最后的一个1
			ret = append(ret, append(arr, 1))
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerate(t *testing.T) {
	fmt.Println(generate(0))
}
