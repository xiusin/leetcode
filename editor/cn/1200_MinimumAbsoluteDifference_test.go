package cn

import (
	"fmt"
	"sort"
	"testing"
)

//给你个整数数组 arr，其中每个元素都 不相同。 
//
// 请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。 
//
// 
//
// 示例 1： 
//
// 输入：arr = [4,2,1,3]
//输出：[[1,2],[2,3],[3,4]]
// 
//
// 示例 2： 
//
// 输入：arr = [1,3,6,10,15]
//输出：[[1,3]]
// 
//
// 示例 3： 
//
// 输入：arr = [3,8,-10,23,19,-4,-14,27]
//输出：[[-14,-10],[19,23],[23,27]]
// 
//
// 
//
// 提示： 
//
// 
// 2 <= arr.length <= 10^5 
// -10^6 <= arr[i] <= 10^6 
// 
// Related Topics 数组
//执行耗时:92 ms,击败了100.00% 的Go用户
//内存消耗:7.5 MB,击败了100.00% 的Go用户
//leetcode submit region begin(Prohibit modification and deletion)
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr) // 先排序数组
	var counter = map[int][][]int{}
	l := len(arr)
	minAbsSubVal := arr[1] - arr[0]
	if minAbsSubVal < 0 {
		minAbsSubVal = -minAbsSubVal
	}
	counter[minAbsSubVal] = [][]int{{arr[0], arr[1]}}
	for i := 2; i < l; i++ {
		sub := arr[i] - arr[i-1]
		if sub < 0 {
			sub = -sub
		}
		if minAbsSubVal > sub {
			delete(counter, minAbsSubVal)
			minAbsSubVal = sub
			counter[minAbsSubVal] = [][]int{{arr[i-1], arr[i]}}
		} else if minAbsSubVal == sub {
			counter[minAbsSubVal] = append(counter[minAbsSubVal], []int{arr[i-1], arr[i]})
		}
	}
	return counter[minAbsSubVal]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumAbsoluteDifference(t *testing.T) {
	fmt.Println(minimumAbsDifference([]int{1,3,6,10,15}))
}
