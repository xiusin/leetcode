package cn

import (
	"fmt"
	"testing"
)

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
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
 //TODO 脑细胞死完了, 被别人两句话搞定
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 递归版
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// l1值小于l2
	if l1.Val < l2.Val {
		// 替换l1的值, 将当前l1节点和l2比较
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}


 // 自写,  每次写这种排序的鬼东西他妈的自己就跟个白痴一样!
func mergeTwoListsSelf(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {	// 异常数据传入判断
		return l2
	} else if l2 == nil {
		return l1
	}
	lastNode, rest := &ListNode{}, &ListNode{} //记录链表最后一个节点
	lastNode = rest                            // 将lastNode指向链表开始位置
	for {
		if l1.Val > l2.Val {
			lastNode.Val = l2.Val //赋值最后结果链表最后一个节点的值
			l2 = l2.Next          // 指向当前链表的下一个节点
		} else {
			lastNode.Val = l1.Val
			l1 = l1.Next
		}
		if l1 == nil || l2 == nil {
			break
		}
		lastNode.Next = &ListNode{}
		lastNode = lastNode.Next
	}
	if l1 != nil {
		lastNode.Next = l1
	} else {
		lastNode.Next = l2
	}
	return rest
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeTwoSortedLists(t *testing.T) {
	list := mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		//nil,
		//nil,
	)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
