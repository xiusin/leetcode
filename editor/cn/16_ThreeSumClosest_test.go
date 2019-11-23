package cn

import (
	"math"
	"sort"
	"testing"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。 
//
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
// 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
//TODO 最接近的三数之和: 还是没有做出来, 差特么一点点!!!
func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return 0
	}
	// 初始化一个数字
	min := nums[0] + nums[1] + nums[l-1] // 初始化第一次迭代的结果, 其实是随意的
	// 求绝对差值
	for k := 0; k < l; k++ {
		i, j := k+1, l-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if math.Abs(float64(target-min)) > math.Abs(float64(target-sum)) {
				min = sum
			} else if sum > target {
				j--
			} else if sum < target {
				i++
			} else {
				return min
			}
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSumClosest(t *testing.T) {
	//println(threeSumClosest([]int{-1, 2, 1, -4}, 1) , 2)
	println(threeSumClosest([]int{-2, 1, 3, 5, 6, 7, 9}, 2), 2)
}
