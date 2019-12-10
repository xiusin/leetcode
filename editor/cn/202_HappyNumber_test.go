package cn

import (
	"fmt"
	"strconv"
	"testing"
)

//编写一个算法来判断一个数是不是“快乐数”。
//
// 一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。 
//
// 示例: 
//
// 输入: 19
//输出: true
//解释: 
//1'2 + 9'2 = 82
//8'2 + 2'2 = 68
//6'2 + 8'2 = 100
//1'2 + 0'2 + 0'2 = 1
//
// Related Topics 哈希表 数学

//leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 要判断循环, 就需要决定是否出现过类似的数字, 比如0的平方
	var ok bool
	var hash = map[int]struct{}{}
	for {
		numStr, tmp := strconv.Itoa(n), 0
		for _, num := range numStr {
			intNum, _ := strconv.Atoi(string(num))
			tmp += intNum * intNum
		}
		if tmp == 1 {
			return true
		}
		if _, ok = hash[tmp]; ok {
			return false
		}
		hash[tmp] = struct{}{}
		n = tmp
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHappyNumber(t *testing.T) {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(0))
	fmt.Println(isHappy(3))
}
