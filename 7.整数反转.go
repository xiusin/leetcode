
/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	flag := false
	if x < 0 {
		flag, x = true, -x
	}
	x, _ = strconv.Atoi(string(reverseString(strconv.Itoa(x))))
	if flag {
		x = -x
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return x
}

func reverseString(x string) (r []byte) {
	for i := len(x) - 1; i >= 0; i-- {
		r = append(r, x[i])
	}
	return
}
