package cn

import (
	"fmt"
	"math"
	"testing"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。 
//
// 示例 1: 
//
// 输入: 123
//输出: 321
// 
//
// 示例 2: 
//
// 输入: -123
//输出: -321
// 
//
// 示例 3: 
//
// 输入: 120
//输出: 21
// 
//
// 注意: 
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。 
// Related Topics 数学

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	// 从各位数逐级向上取余
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	var rest int
	var i = 1
	var list []int
	for x > 0 {
		list = append(list, x%10)
		x = (x - x%10) / 10	//逐个剔除末位并降低数量级
		i++
	}
	lens, util := len(list),1
	for i = lens - 1; i>=0; i-- {
		rest += util * list[i]
		util *= 10	//递增数量级
	}
	if neg {
		rest = -rest
	}
	if rest < math.MinInt32 || rest > math.MaxInt32 {
		return 0
	}
	return rest
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseInteger(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(reverse(3))
	fmt.Println(reverse(2))
	fmt.Println(reverse(0))
}
