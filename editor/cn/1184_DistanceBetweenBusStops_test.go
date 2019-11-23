package cn

import (
	"fmt"
	"testing"
)

//环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。 
//
// 环线上的公交车都可以按顺时针和逆时针的方向行驶。 
//
// 返回乘客从出发点 start 到目的地 destination 之间的最短距离。 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：distance = [1,2,3,4], start = 0, destination = 1
//输出：1
//解释：公交站 0 和 1 之间的距离是 1 或 9，最小值是 1。 
//
// 
//
// 示例 2： 
//
// 
//
// 输入：distance = [1,2,3,4], start = 0, destination = 2
//输出：3
//解释：公交站 0 和 2 之间的距离是 3 或 7，最小值是 3。
// 
//
// 
//
// 示例 3： 
//
// 
//
// 输入：distance = [1,2,3,4], start = 0, destination = 3
//输出：4
//解释：公交站 0 和 3 之间的距离是 6 或 4，最小值是 4。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10^4 
// distance.length == n 
// 0 <= start, destination < n 
// 0 <= distance[i] <= 10^4 
// 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	// 判断交换顺序
	if start > destination {
		start, destination = destination, start
	}
	// 求总和
	//totalDistance := (1 + n) * n / 2 // 高斯算法 根据测试用例发现间距非递增等宽
	//startToDest := (distance[destination - 1] + distance[start]) * (destination - 1 - start) / 2
	totalDistance := 0
	startToDest := 0
	for k, v := range distance {
		totalDistance += v
		if k >= start && k < destination { // 不包含destination位置的数字,因为是此位置与下站的距离
			startToDest += v
		}
	}
	destToStart := totalDistance - startToDest
	// 笨方法
	if startToDest > destToStart {
		return destToStart
	} else {
		return startToDest
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistanceBetweenBusStops(t *testing.T) {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1) == 1)
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2) == 3)
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3) == 4)
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 1, 3) == 5)
	fmt.Println(distanceBetweenBusStops([]int{7, 6, 3, 0, 3}, 0, 4) == 3)

}
