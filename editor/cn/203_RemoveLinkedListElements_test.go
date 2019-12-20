package cn

import (
	"encoding/json"
	"fmt"
	"testing"
)

//删除链表中等于给定值 val 的所有节点。 
//
// 示例: 
//
// 输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
// 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	var ret = &ListNode{Next:head}	//使用一个节点后缀上链表
	var prev = ret	// 使用临时节点保存当前遍历的节点
	// 找到第一个不为目标值的节点
	for prev.Next != nil {
		if prev.Next.Val != val {
			prev = prev.Next
		} else {
			prev.Next = prev.Next.Next
		}
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveLinkedListElements(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}}}
	s, _ := json.Marshal(removeElements(l, 6))
	fmt.Printf(string(s))
}
