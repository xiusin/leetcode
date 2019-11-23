package cn

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	l := len(nums)
	i, j := 0, l-1
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			break
		}
	}
	return []int{i, j}
}

// 最蠢的算法
//func versionOne(nums []int, target int) int{
//	var res []int
	// 记录数据所在的索引, 假设存在相同的值
	//ss, l := map[int][]int{}, len(nums)
	//if l > 0 {
	//	hp, ep := 0, l-1
	//	for idx, num := range nums {
	//		ss[num] = append(ss[num], idx)
	//	}
	//	sort.Slice(nums, func(i, j int) bool {
	//		return nums[i] < nums[j]
	//	})
	//	for hp < ep && nums[hp]+nums[ep] != target {
	//		if nums[hp]+nums[ep] < target {
	//			hp++
	//		} else if nums[hp]+nums[ep] > target {
	//			ep--
	//		}
	//	}
	//	if nums[hp]+nums[ep] == target {
	//		idx1 := ss[nums[hp]]
	//		if nums[hp] != nums[ep] {
	//			idx2 := ss[nums[ep]]
	//			return append(res, idx1[0], idx2[0])
	//		} else {
	//			return append(res, idx1[0], idx1[1])
	//		}
	//	}
	//}
//}

//leetcode submit region end(Prohibit modification and deletion)
func TestTwoSum(t *testing.T) {
	tests := []struct {
		data   []int
		target int
		expect []int
	}{
		//{data: []int{1, 2, 4, 3, 5}, target: 8, expect: []int{3, 4}},
		//{data: []int{3, 2, 2, 5}, target: 4, expect: []int{1, 2}},
		{data: []int{230, 863, 916, 585, 981, 404, 316, 785, 88, 12, 70, 435, 384, 778, 887, 755, 740, 337, 86, 92, 325, 422, 815, 650, 920, 125, 277, 336, 221, 847, 168, 23, 677, 61, 400, 136, 874, 363, 394, 199, 863, 997, 794, 587, 124, 321, 212, 957, 764, 173, 314, 422, 927, 783, 930, 282, 306, 506, 44, 926, 691, 568, 68, 730, 933, 737, 531, 180, 414, 751, 28, 546, 60, 371, 493, 370, 527, 387, 43, 541, 13, 457, 328, 227, 652, 365, 430, 803, 59, 858, 538, 427, 583, 368, 375, 173, 809, 896, 370, 789}, expect: []int{28, 45}, target: 542},
	}
	for k, test := range tests {
		fmt.Println(test.data[test.expect[0]], test.data[test.expect[1]])
		if ret := twoSum(test.data, test.target); reflect.DeepEqual(ret, test.expect) {
			t.Log("测试第", k, "组数据成功")
		} else {
			t.Log("测试第", k, "组数据失败", "结果: ", ret)
		}
	}
}
