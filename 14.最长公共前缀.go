package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.31%)
 * Likes:    634
 * Dislikes: 0
 * Total Accepted:    104.4K
 * Total Submissions: 304.1K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
	arrlen := len(strs)
	if arrlen == 0 { // 为空直接返回
		return ""
	} else if arrlen == 1 { // 只有一个元素直接返回
		return strs[0]
	}
	strlen, currentPrefix := len(strs[0]), ""
	if strlen != 0 {
		for i := 0; i < strlen; i++ {
			prefix := strs[0][:i+1]
			// 修改
			for j := 1; j < arrlen; j++ {
				// 匹配字符串长度不得溢出索引,并且所有的元素均匹配通过
				if len(strs[j]) <= i || strs[j][:i+1] != prefix {
					return currentPrefix
				}
			}
			currentPrefix = prefix
		}
	}
	return currentPrefix
}

// 水平扫描
func longestCommonPrefix2(strs []string) string {
	arrlen := len(strs)
	if arrlen == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ { // 迭代字符串长度
		char := strs[0][i : i+1] // 得到每个位置的字符
		for j := 1; j < arrlen; j++ {
			// 如果当前位置的索引值与其中一个元素的总长度一致，返回
			// 如果当前索引位置的字符不一致则返回
			if i == len(strs[j]) || strs[j][i:i+1] != char {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

//https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode/
// 水平扫描法
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 依次往后对比， 如S1和S2查出来最长公串，再与S3。。对比找出最长公串
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// 分治法
// 二分查找法
// 前缀树
