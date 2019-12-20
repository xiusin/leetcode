package cn

import (
	"fmt"
	"testing"
)

//我们正在玩一个猜数字游戏。 游戏规则如下： 
//我从 1 到 n 选择一个数字。 你需要猜我选择了哪个数字。 
//每次你猜错了，我会告诉你这个数字是大了还是小了。 
//你调用一个预先定义好的接口 guess(int num)，它会返回 3 个可能的结果（-1，1 或 0）： 
//
// -1 : 我的数字比较小
// 1 : 我的数字比较大
// 0 : 恭喜！你猜对了！
// 
//
// 示例 : 
//
// 输入: n = 10, pick = 6
//输出: 6 
// Related Topics 二分查找
func guess(n int) int {
	if GuessNumber > n {
		return -1
	} else if GuessNumber < n {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right - left) / 2
		switch guess(mid) {
		case -1:	// 如果猜小了, 就移动左边位置
			left = mid + 1
		case 1:
			right = mid - 1
		default:
			return mid
		}
	}
	return left
}

const GuessNumber  = 6
func TestGuessNumberHigherOrLower(t *testing.T)  {
	fmt.Println(guessNumber(10))
	fmt.Println(guessNumber(23))
	fmt.Println(guessNumber(1000))
}