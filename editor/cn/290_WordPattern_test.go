package cn

import (
	"fmt"
	"strings"
	"testing"
)

//给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//
// 示例1:
//
// 输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//
// 示例 2:
//
// 输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//
// 示例 3:
//
// 输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//
// 示例 4:
//
// 输入: pattern = "abba", str = "dog dog dog dog"
//输出: false
//
// 说明:
//你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
// Related Topics 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	var patternWordMaps = map[byte]string{} // 记录pattern和word的关系
	var wordPatternMaps = map[string]byte{} // 记录word与pattern字母的关系
	wl := len(words)
	var res []byte
	if wl == len(pattern) {
		for i := 0; i < wl; i++ {
			// 匹配模式
			if word, ok := patternWordMaps[pattern[i]]; !ok { // 判断字母是否已经映射过单词 (如果只有单一的这个判断可能出现 a => dog, b => dog)
				if _, ok1 := wordPatternMaps[words[i]]; !ok1 { // 过滤多个字母对应同一个单词的场景
					patternWordMaps[pattern[i]] = words[i]
					wordPatternMaps[words[i]] = pattern[i]
					res = append(res, pattern[i])
				}
			} else if word == words[i] { // 存在的并且字典里保存到值等于当前word
				res = append(res, wordPatternMaps[word])
			}
		}
		if string(res) == pattern {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordPattern(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog") == true)
	fmt.Println(wordPattern("abba", "dog cat cat fish") == false)
	fmt.Println(wordPattern("aaaa", "dog cat cat dog") == false)
	fmt.Println(wordPattern("aaaa", "dog dog dog dog") == true)
	fmt.Println(wordPattern("abba", "dog dog dog dog") == false)
	fmt.Println(wordPattern("ab", "dog dog") == false)
}
