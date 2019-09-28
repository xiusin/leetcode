package cn

import (
	"testing"
)

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
// 案例: 
//
// 
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
// 
//
// 
//
// 注意事项：您可以假定该字符串只包含小写字母。 
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
//解答成功:
//执行耗时:12 ms,击败了88.51% 的Go用户
//内存消耗:5.7 MB,击败了90.06% 的Go用户
func firstUniqChar(s string) int {
	m, cs := [26][2]int{}, []byte{}
	for idx := range s {
		c := s[idx] - 'a'	// 基于字符a向前位移索引到26以内
		if val := m[c]; val[0] == 0 {
			m[c][0] = 1
			m[c][1] = idx
			cs = append(cs, c)
		} else {
			m[c][0]++
		}
	}
	for _, c := range cs {
		if val := m[c]; val[0] == 1 {
			return val[1]
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstUniqueCharacterInAString(t *testing.T) {

	tests := []struct {
		s      string
		expect int
	}{
		//{"leetcode", 0},
		//{"loveleetcode", 2},
		{"z", 0},
	}
	for k, test := range tests {
		if ret := firstUniqChar(test.s); ret == test.expect {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败")
		}
	}

}
