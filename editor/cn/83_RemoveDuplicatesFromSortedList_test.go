package cn

import (
	"encoding/json"
	"fmt"
	"testing"
)

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。 
//
// 示例 1: 
//
// 输入: 1->1->2
//输出: 1->2
// 
//
// 示例 2: 
//
// 输入: 1->1->2->3->3
//输出: 1->2->3 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var curr = head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// 调整next的节点
			curr.Next = curr.Next.Next
		} else {
			// 更换curr节点
			curr = curr.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	ListNodes := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
		},
	}
	v, _ := json.Marshal(deleteDuplicates(ListNodes))
	fmt.Println(string(v))
}
