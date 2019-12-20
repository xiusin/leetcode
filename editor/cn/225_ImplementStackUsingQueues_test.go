package cn

import (
	"fmt"
	"testing"
)

//使用队列实现栈的下列操作： 
//
// 
// push(x) -- 元素 x 入栈 
// pop() -- 移除栈顶元素 
// top() -- 获取栈顶元素 
// empty() -- 返回栈是否为空 
// 
//
// 注意: 
//
// 
// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。 
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。 
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。 
// 
// Related Topics 栈 设计

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	_data []int
	len   int
}

/** Initialize your data structure here. */
func StackConstructor() MyStack {
	return MyStack{_data: make([]int, 0), len: 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this._data = append([]int{x}, this._data...)
	this.len++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.len > 0 {
		this.len--
		el := this._data[0]
		this._data = this._data[1:]
		return el
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.len > 0 {
		return this._data[0]
	} else {
		return 0
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.len == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementStackUsingQueues(t *testing.T) {
	obj := StackConstructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top(), obj.Pop(), obj.Empty())
}
