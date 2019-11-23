package cn

import (
	"fmt"
	"testing"
)

//在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。
// 给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
//
// 返回可以通过分割得到的平衡字符串的最大数量。
// 示例 1：
// 输入：s = "RLRRLLRLRL"
//输出：4
//解释：s 可以分割为 "RL", "RRLL", "RL", "RL", 每个子字符串中都包含相同数量的 'L' 和 'R'。
// 示例 2：
// 输入：s = "RLLLLRRRLR"
//输出：3
//解释：s 可以分割为 "RL", "LLLRRR", "LR", 每个子字符串中都包含相同数量的 'L' 和 'R'。
// 示例 3：
//
// 输入：s = "LLLLRRRR"
//输出：1
//解释：s 只能保持原样 "LLLLRRRR".
// 提示：
// 1 <= s.length <= 1000
// s[i] = 'L' 或 'R'
//
// Related Topics 贪心算法 字符串

func balancedStringSplit(s string) int {
	//return MyErrorDemo(s)
	// 遍历字符串,  首先L和R相等.
	var res, num int
	for _, v := range s {
		if v == 'L' {
			num++
		} else {
			num--
		}
		if num == 0 {
			res++
		}
	}
	return res


}

//leetcode submit region end(Prohibit modification and deletion)
// 之前题意看成了: 随机打乱, 只不过是非绝对的由几段平衡字符串组成. 如: RRLLLLRLRR 这类的非绝对平衡字符串完成
// 其实测试用例是绝对平衡字符串, 从头开始分割绝对都能用上所有的字符.
func MyErrorDemo(s string) int {
	var lNum, rNum, res int
	// 在L有值直到碰到R之前可以在累加状态然后直到下次碰到L或等于之前累积的L的值, 相反, 如果不对等, 就踢掉先累积的值
	var curBalanceStartChar int32
	var locker bool
	for _, char := range s {
		if curBalanceStartChar == 0 {
			curBalanceStartChar = char 	// 记录小平衡串判断的起点
		}
		if char == curBalanceStartChar && !locker {
			lNum++
		} else if char != curBalanceStartChar {
			rNum++
			locker = true	// 只要碰到不是起点的字符就锁定起点字符的累积
		}
		// 左边等于右边时 可以匹配左边小右边大的情况
		if lNum == rNum {
			res++
			curBalanceStartChar, lNum, rNum, locker = 0, 0, 0, false
		}
	}
	// 处理最后匹配
	if lNum > rNum {
		res++
	}
	return res
}

func TestSplitAStringInBalancedStrings(t *testing.T) {
	fmt.Println(balancedStringSplit("RLRRRLLRLL") == 2)
	fmt.Println(balancedStringSplit("LLLLRRRR") == 1)
	fmt.Println(balancedStringSplit("RLLLLRRRLR") == 3)
	fmt.Println(balancedStringSplit("RLRRLLRLRL") == 4)
}
