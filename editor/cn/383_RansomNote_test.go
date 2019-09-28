package cn

import (
	"testing"
)

//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//
// (题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。) 
//
// 注意： 
//
// 你可以假设两个字符串均只含有小写字母。 
//
// 
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true
// 
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	val, allMagazineChar := 0, make(map[byte]int, 26)
	rl, ml := len(ransomNote), len(magazine)
	for i := 0; i < ml; i++ {
		allMagazineChar[magazine[i]] += 1 // map的元素可以自动初始化
	}
	for i := 0; i < rl; i++ {
		if val = allMagazineChar[ransomNote[i]]; val == 0 {
			return false
		} else {
			allMagazineChar[ransomNote[i]]--
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRansomNote(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expect     bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for k, test := range tests {
		if ret := canConstruct(test.ransomNote, test.magazine); ret == test.expect {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败")
		}
	}
}
