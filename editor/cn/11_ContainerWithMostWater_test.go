package cn

import (
	"fmt"
	"testing"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。 
//
// 
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 
//
// 示例: 
//
// 输入: [1,8,6,2,5,4,8,3,7]
//输出: 49 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	var maxArea int
	var area, head, tail = 0, 0, len(height)-1
	for head < tail {
		width := tail - head
		if height[tail] > height[head] {
			area = width * height[head]
			head++ //下次指针移动高度较小的指针移动(因为面积受最小高度限制, 除非遇到了比这个更大的值才有增大面积的可能性)
		} else {
			area = width * height[tail]
			tail--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
