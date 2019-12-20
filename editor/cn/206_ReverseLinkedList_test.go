package cn

import (
	"encoding/json"
	"fmt"
	"testing"
)

//反转一个单链表。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
// todo 翻转链表 不会做, 先抄的题解
func reverseList(head *ListNode) *ListNode {
	// 双指针迭代
	// reverseLink 最终结果
	// cur 迭代的当前指针
	var reverseLink *ListNode
	var cur = head
	for cur != nil {
		// 记录当前节点的下一个节点
		tmp := cur.Next
		// 将当前节点指向reverseLink
		cur.Next = reverseLink
		//reverseLink和cur节点都向后一位
		reverseLink = cur
		cur = tmp
	}
	return reverseLink
}

// 递归处理
func TestReverseLinkedList(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	s, _ := json.Marshal(reverseList(l))
	fmt.Printf(string(s))
}
