package cn

import (
	"fmt"
	"strings"
	"testing"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。 
//
// 说明：本题中，我们将空字符串定义为有效的回文串。 
//
// 示例 1: 
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "race a car"
//输出: false
// 
// Related Topics 双指针 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	var i, j = 0, len(s)
	if j >= 2 {
		j--
		s = strings.ToLower(s)
		for i < j {
			// 只验证 a-z A-Z 0-9
			if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9')) {
				i++
			} else if !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9')) {
				j--
			} else if s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("op"))
	fmt.Println(isPalindrome("ab"))
	fmt.Println(isPalindrome("ab11ba"))
	fmt.Println(isPalindrome("ab12a"))
}
