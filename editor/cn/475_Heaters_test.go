package cn

import (
	"sort"
	"testing"
)

//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
// 现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
//
// 所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。
//
// 说明:
//
//
// 给出的房屋和供暖器的数目是非负数且不会超过 25000。
// 给出的房屋和供暖器的位置均是非负数且不会超过10^9。
// 只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
// 所有供暖器都遵循你的半径标准，加热的半径也一样。
//
//
// 示例 1:
//
//
//输入: [1,2,3],[2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//
//
// 示例 2:
//
//
//输入: [1,2,3,4],[1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//
// Related Topics 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
func findRadius(houses []int, heaters []int) int {
	/**
	解题过程:
	对每个房子找到离它【最近】的供暖器；
	找到 房子-供暖器距离 的【最大】值
	*/
	sort.Ints(houses)
	sort.Ints(heaters)

	return 0
	//if housesTotal > 0 && heatersTotal > 0 {
	//	maxDistance = int(math.Abs(float64(houses[0] - heaters[0])))
	//	var currHousePos int
	//	var curDistance int
	//	for k, currentHeater := range heaters {
	//		if k == 0 {
	//			curDistance = int(math.Abs(float64(currentHeater - houses[0]))) //计算开始供暖器位置离第一个房子的距离
	//		} else if currHousePos < housesTotal {
	//			curDistance = (currentHeater - houses[currHousePos-1]) / 2 // 计算两个供暖期位置最大的辐射半径
	//		}
	//		if maxDistance < curDistance {
	//			maxDistance = curDistance
	//		}
	//		currHousePos = currentHeater
	//	}
	//}
	//if housesTotal > 1 && heaters[heatersTotal-1] < houses[housesTotal-1] {
	//	endDistance := houses[housesTotal-1] - heaters[heatersTotal-1]
	//	fmt.Println(endDistance)
	//	if maxDistance < endDistance {
	//		maxDistance = endDistance
	//	}
	//}
	//return maxDistance
}

//leetcode submit region end(Prohibit modification and deletion)
func TestHeaters(t *testing.T) {
	tests := []struct {
		houses  []int
		heaters []int
		expect  int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1, 2, 3, 4}, []int{1}, 3},
		{[]int{1, 2, 3, 4}, []int{2}, 2},
		{[]int{1, 5}, []int{2}, 3},
		{[]int{2, 7, 88}, []int{7}, 81},
		{[]int{1, 3, 5}, []int{3}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 4}, 3},
		{[]int{1}, []int{3, 4}, 2},
		{[]int{}, []int{3, 4}, 0},
		{[]int{5, 4, 3, 2, 1}, []int{2}, 3},
		{[]int{1, 2, 3, 5, 15}, []int{2, 30}, 13},
	}
	for k, test := range tests {
		if ret := findRadius(test.houses, test.heaters); ret == test.expect {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败, 期望", test.expect, "实际", ret)
		}
	}
}
