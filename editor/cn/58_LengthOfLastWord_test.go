package cn

import (
	"fmt"
	"testing"
)

//给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。 
//
// 如果不存在最后一个单词，请返回 0 。 
//
// 说明：一个单词是指由字母组成，但不包含任何空格的字符串。 
//
// 示例: 
//
// 输入: "Hello World"
//输出: 5
// 
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	// 判断最后一个单词
	//subStrs := strings.Split(s, " ")
	//j, l := 0, len(subStrs)
	//for i := l - 1; i >= 0; i-- {
	//	j = 0
	//	for k := range subStrs[i] {
	//		v := subStrs[i][k]
	//		if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) {
	//			j = 0
	//			break
	//		} else {
	//			j++
	//		}
	//	}
	//	if j > 0 {
	//		return j
	//	}
	//}

	counter, l := 0, len(s)
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if counter != 0 {
				return counter
			}
			counter = 0
		} else if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			for i >= 0 && s[i] != ' ' {
				i--
			}
			counter = 0
		} else {
			counter++
		}
	}
	return counter
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("1asdasd d 43sdfasd"))
}
