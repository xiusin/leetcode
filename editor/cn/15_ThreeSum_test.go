package cn

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	sort.Ints(nums)
	if l >= 3 && nums[l-1] >= 0 {
		for k := 0; k < l-2; k++ {
			v := nums[k]
			//TODO 三数之和: 过滤相同值
			if k > 0 && nums[k -1] == nums[k] {	// 这里过滤连续两个相同的值
				continue
			}
			if v <= 0 {
				var i, j = k+1, l-1
				for i < j {
					sum := v + nums[i] + nums[j]
					if sum == 0 {
						for i < j && nums[i] == nums[i+1] {	// 搜索结果集的时候同理, 过滤相同的值
							i++
						}
						for i < j && nums[j] == nums[j-1] {
							j--
						}
						res = append(res, []int{v, nums[i], nums[j]}) // 此时如果再移动值要么是跟此值重复, 要么就sum值一直增大. (i最好也就保持不变, j最好也是保持不变)
						i++
						j--
					} else if sum > 0 {
						j--
					} else {
						i++
					}
				}
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	//fmt.Println(threeSum([]int{-1, 3, -3, -2, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-4, -1, -1, -1, 0, 1, 2,}))
	fmt.Println(threeSum([]int{-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0}))
	//fmt.Println(threeSum([]int{-1, 3}))
	//fmt.Println(threeSum([]int{0, 0, 0, 0}))
	//fmt.Println(threeSum([]int{-1, -1, 0, 1}))
	//fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
