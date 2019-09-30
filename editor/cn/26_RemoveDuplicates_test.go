package cn

import (
	"testing"
)

//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//示例 1:
//
//给定数组 nums = [1,1,2],
//
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//
//你不需要考虑数组中超出新长度后面的元素。
//说明:
//
//为什么返回数值是整数，但输出的答案是数组呢?
//
//请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
//你可以想象内部操作如下:
//
//nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
//int len = removeDuplicates(nums);
//
//在函数里修改输入数组对于调用者是可见的。
//根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}

//https://leetcode-cn.com/problems/remove-element/solution/yi-chu-yuan-su-by-leetcode/
//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	// todo 不会做
	/**
	(同向指针, 快慢的却别)
	1. 测试用例: 1, 1, 1, 2, 3, 4, 4
	2. 初始化i指针指向0的位置, 初始化j指针指向1的位置
	4. i不动, 移动j的指针, 直到找到不相同的位置
	5. 后移i的指针, 为什么一定能移动呢? (试想一下, 当指针i和j是相邻的位置并且值不同, i =1, j=2, 当i变成j 其实就相当与i = j 属于自身交换值, 没有改变)
	6. 根据本例, 最终变成了 1, 2 (i指向的位置), 1,2(j指向的位置),3,4,4
	7. 继续迭代 j向后移动, 指向了3 , 此时i的值和j的值不同, 则后移i的位置覆盖数据: 1, 2 , 3(i指向的位置),2,3(j指向的位置),4,4
	8. 继续迭代 j向后移动, 指向了4 , 此时i的值和j的值不同, 则后移i的位置覆盖数据: 1, 2 , 3,4,(i指向的位置),3,4(j指向的位置),4
	9. 继续迭代 j向后移动, 指向了5 , 此时i的值和j的值相同
	10. 结束, 此时i的值为最终的索引位置, 还需要计算为总数,即 i+1
	*/

	i := 0 // 索引 or 计数  慢指针
	if len(nums) > 0 {
		for j := 1; j < len(nums); j++ {
			if nums[j] != nums[i] {
				// 后移指针覆盖数据
				i++
				// 使用j 覆盖 i
				nums[i] = nums[j]
			}
		}
		i = i + 1	// 最终索引位置+1 = 总数
	}
	return i
}
//leetcode submit region end(Prohibit modification and deletion)
func TestRemoveDuplicates(t *testing.T) {
	// 提交时的错误时的测试用例
	tests := []struct {
		data   []int
		except int
	}{
		{data: []int{1, 1, 1, 2, 3, 4, 4}, except: 4},
		{data: []int{1, 1, 1, 1}, except: 1}, {data: []int{}, except: 0},
		{data: []int{1, 2}, except: 2},
	}
	for k, test := range tests {
		if ret := removeDuplicates(test.data); ret == test.except {
			t.Log("测试第", k, "组数据成功", test.data)
		} else {
			t.Error("测试第", k, "组数据失败", test.data)
		}
	}
}
