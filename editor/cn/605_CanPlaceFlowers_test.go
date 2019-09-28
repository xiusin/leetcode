package cn

import "testing"

//假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。
//
// 示例 1: 
//
// 
//输入: flowerbed = [1,0,0,0,1], n = 1
//输出: True
// 
//
// 示例 2: 
//
// 
//输入: flowerbed = [1,0,0,0,1], n = 2
//输出: False
// 
//
// 注意: 
//
// 
// 数组内已种好的花不会违反种植规则。 
// 输入的数组长度范围为 [1, 20000]。 
// n 是非负整数，且不会超过输入数组的大小。 
// 
// Related Topics 数组
//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n > 0 { // 记录连续的空地, 初始化空地为1, 向前补一位(具体为啥 先做)
		emptyFlowered ,l  := 1, len(flowerbed)
		var canPlaceFirst, odd int
		for i :=0 ; i< l; i++ {
			if flowerbed[i] == 1 && emptyFlowered != 0 {
				// 检查空地, 计算可以种的花朵, 此处是夹心饼干的情况
				canPlaceFirst,odd = emptyFlowered/2, emptyFlowered%2
				if odd == 0 && canPlaceFirst-1 > 0 {
					n = n - (canPlaceFirst - 1)
				} else if odd == 1 && canPlaceFirst > 0 {
					n = n - canPlaceFirst
				}
				emptyFlowered = 0 // 碰到1 则清空
				if n <= 0 {
					break
				}
			} else {
				emptyFlowered = emptyFlowered + 1 // 空地计入
			}
		}
		// 如果还有空地(此时为连续空地左边为1 右边到头的情况)
		if n > 0 && emptyFlowered > 0 {
			n = n - emptyFlowered/2 // 100 => 1, 1000 => 1, 10000 => 2 推算
		}
	}
	if n > 0 {
		return false
	} else {
		return true
	}
}
//解答成功:
//执行耗时:20 ms,击败了99.22% 的Go用户
//内存消耗:5.9 MB,击败了27.27% 的Go用户  todo 为什么内存占用这么高
//leetcode submit region end(Prohibit modification and deletion)
func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		data []int
		n    int
		ex   bool
	}{
		{[]int{1}, 1, false},
		{[]int{0}, 1, true},
		{[]int{1, 1}, 1, false},
		{[]int{1, 0, 0}, 1, true},
		{[]int{0, 1, 0}, 1, false},
		{[]int{1, 0, 0}, 2, false},
		{[]int{1, 0, 0, 0}, 2, false},
		{[]int{1, 0, 0, 0, 0}, 2, true},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 3, false},
	}
	for k, test := range tests {
		if ret := canPlaceFlowers(test.data, test.n); ret == test.ex {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败")
		}
	}
}
