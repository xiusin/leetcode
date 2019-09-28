package cn

import (
	"strconv"
	"testing"
)

//实现 strStr() 函数。
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。
// 示例 1:
// 输入: haystack = "hello", needle = "ll"
//输出: 2
// 示例 2:
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
// 说明:
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
// Related Topics 双指针 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	haystackLen, needleLen := len(haystack), len(needle)
	if haystackLen >= needleLen {
		//haystackLen - needleLen + 1
		// 比如 hello 和 ll , 最小迭代字符串为: hell
		// 12 和 1, 最小迭代字符串 12
		for i := 0; i < haystackLen-needleLen+1; i++ {
			if haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}
	return -1
}

//https://leetcode-cn.com/problems/implement-strstr/solution/kmp-suan-fa-xiang-jie-by-labuladong/
// todo KMP算法详解
type KMP struct {
	dp  [][256]int // 牛逼叫法: 有限状态机
	pat string  //匹配符或待匹配字符串
	M int
}

func NewKMP(pat string) *KMP {
	kmp := &KMP{pat: pat}
	kmp.M = len(pat)
	if kmp.M > 0 {
		// dp[状态][字符] = 下一个状态
		for i := 0; i<kmp.M; i++ {
			kmp.dp = append(kmp.dp, [256]int{})
		}
		//初始化状态
		kmp.dp[0][kmp.pat[0]] = 1
		// 影子状态: 所谓影子状态，就是和当前状态具有相同的前缀
		X := 0
		// 当前状态J从 1开始
		for j := 1; j<kmp.M; j++  {
			for c := 0; c < 256; c++ {
				if int(byte(pat[j])) == c {
					kmp.dp[j][c] = j+1
				} else {
					kmp.dp[j][c] = kmp.dp[X][c]
				}
			}
			// 更新影子状态
			X = kmp.dp[X][kmp.pat[j]]
		}
	}
	return kmp
}

func (k *KMP) search(txt string) int {
	 N :=  len(txt)
	 if k.M == 0 {
		 return 0
	 }
	if k.M > N {
		return -1
	}
	// pat初始化状态为0
	j := 0
	for i := 0; i < N; i++ {
		// 当前是状态 j，遇到字符 txt[i]，
		// pat 应该转移到哪个状态？
		j = k.dp[j][int(byte(txt[i]))]
		// 如果达到终止态，返回匹配开头的索引
		if j == k.M {
			return i - k.M + 1
		}
	}
	return -1
}

var ts = [][]string{
	{"hello", "ll", "2"},
	{"", "1", "-1"},
	{"", "", "0"},
	{"12", "2", "1"},
	{"122", "23", "-1"},
}

func TestKMP(t *testing.T) {
	for _, v := range ts {
		idx, _ := strconv.Atoi(v[2])
		kmp := NewKMP(v[1])
		if index := kmp.search(v[0]); index == idx {
			t.Logf("haystack: %s    needle: %s    index: %d  success!! \n", v[0], v[1], idx)
		} else {
			t.Errorf("haystack: %s    needle: %s    index: %d , get: %d failed \n", v[0], v[1], idx, index)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestImplementStrstr(t *testing.T) {
	for _, v := range ts {
		idx, _ := strconv.Atoi(v[2])
		if index := strStr(v[0], v[1]); index == idx {
			t.Logf("haystack: %s    needle: %s    index: %d  success!! \n", v[0], v[1], idx)
		} else {
			t.Errorf("haystack: %s    needle: %s    index: %d , get: %d failed \n", v[0], v[1], idx, index)
		}

	}
}
