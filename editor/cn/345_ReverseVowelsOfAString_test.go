package cn

import (
	"fmt"
	"testing"
)

//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
// 示例 1:
//
// 输入: "hello"
//输出: "holle"
//
//
// 示例 2:
//
// 输入: "leetcode"
//输出: "leotcede"
//
// 说明:
//元音字母不包含字母"y"。
// Related Topics 双指针 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	var vowels = map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	sl, ss := len(s), []byte(s)
	if sl >= 2 { // 2个以上
		var i, j = 0, sl - 1
		for i < j {
			_, ok1 := vowels[ss[i]]
			if !ok1 {
				i++
				continue
			}
			_, ok2 := vowels[ss[j]]
			if !ok2 {
				j--
				continue
			}
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
		}
	} else {
		return s
	}
	return string(ss)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseVowelsOfAString(t *testing.T) {
	fmt.Println(reverseVowels("aA"))
}
