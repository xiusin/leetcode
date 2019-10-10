/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.44%)
 * Likes:    1208
 * Dislikes: 0
 * Total Accepted:    77.2K
 * Total Submissions: 329.5K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

import "sort"

// 解题思路：https://leetcode-cn.com/problems/3sum/solution/hua-jie-suan-fa-15-san-shu-zhi-he-by-guanpengchn/
func threeSum(nums []int) (res [][]int) {
	arrlen := len(nums)
	if nums == nil || arrlen < 3 {
		return nil
	}
	// 需要使用双指针法
	// 1. 先排序切片（数组）， 达到有序
	sort.Ints(nums)
	for i := 0; i < arrlen; i++ {
		if nums[i] > 0 {
			break
		} //如果当前数字已经大于零， 那么后面的数字都会大于零
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} //连续两个相同的值跳过
		//2. 选定一个值作为定值
		//3. 选定后一位以及结束位作为双指针向内靠拢
		lIdx, rIdx := i+1, arrlen-1
		for lIdx < rIdx {
			sum := nums[i] + nums[lIdx] + nums[rIdx]
			if sum == 0 {
				// 4. 符合条件的数据添加到一个数组内
				res = append(res, []int{nums[i], nums[lIdx], nums[rIdx]})
				for lIdx < rIdx && nums[lIdx] == nums[lIdx+1] {
					lIdx++ // 过滤去重，处理连续两个相通的值
				}
				for lIdx < rIdx && nums[rIdx] == nums[rIdx-1] {
					rIdx-- // 过滤去重，处理连续两个相通的值
				}
				// 向内移动指针
				lIdx++
				rIdx--
			} else if sum > 0 {
				rIdx-- // 如果值大于零，移动右指针 （移动左指针会继续增大）
			} else if sum < 0 {
				lIdx++ // 如果值小于零，移动左指针 （移动右指针会继续减小）
			}
		}
	}
	return
}
