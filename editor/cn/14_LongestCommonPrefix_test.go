package cn

import (
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	// 找出最短字符串长度 不用处理越界了
	var strMinLength int
	var shortStr string
	if len(strs) > 0 {
		strMinLength, shortStr = len(strs[0]), strs[0]
		for i := range strs {
			if len(strs[i]) < strMinLength {
				strMinLength = len(strs[i])
				shortStr = strs[i]
			}
		}
		if strMinLength > 0 {
			for i := 0; i < strMinLength; i++ {
				for _, str := range strs {
					if str[i] != shortStr[i] {
						return shortStr[:i]
					}
				}
			}
		}
		return shortStr
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		data   []string
		expect string
	}{
		{[]string{"aasdasdasd", "aa"}, "aa"},
		{[]string{"aasdasdasd", ""}, ""},
		{[]string{}, ""},
		{[]string{"abasdjdasdasd", "abassjdasdasd"}, "abas"},
	}
	for k, test := range tests {
		if ret := longestCommonPrefix(test.data); ret == test.expect {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败", "结果: ", ret)
		}
	}
}
