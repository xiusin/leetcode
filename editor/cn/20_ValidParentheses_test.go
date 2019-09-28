package cn

import (
	"testing"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	// 开始和闭合之前不能有单一的其他括号的闭合
	// 双指针貌似不大可行. 栈比较可行
	// 栈低是左括号 栈顶是对于的右括号
	// 最小栈必须符合规则才能向外扩散规则, 即: {} [] ()
	l, salt := len(s), map[byte]byte{'(': ')', '{': '}', '[': ']'}
	if l == 0 {
		return true
	} else if l%2 != 0 {
		return false
	}
	var stack []byte // 数组模拟吧
	var ss int
	for idx := range s {
		if _, ok := salt[s[idx]]; !ok { // 没有找到map 则是右括号
			if ss = len(stack); ss == 0 {
				return false
			}
			// 取出上一个字符, 不匹配返回失败
			if pc := stack[ss-1]; salt[pc] != s[idx] {
				return false
			}
			// 匹配则踢出上一个字符
			stack = stack[0 : ss-1]
		} else {
			stack = append(stack, s[idx])
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {
	// 提交时的错误时的测试用例
	tests := []struct {
		data string
		ex   bool
	}{
		{data: "", ex: true},
		{data: "{[]}", ex: true},
		{data: "{(]}", ex: false},
		{data: "([)]", ex: false},
		{data: "(]", ex: false},
		{data: "([]", ex: false},
		{data: "()[]{}", ex: true},
	}
	for k, test := range tests {
		if ret := isValid(test.data); ret == test.ex {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败")
		}
	}
}
