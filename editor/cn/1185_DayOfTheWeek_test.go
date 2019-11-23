package cn

import (
	"fmt"
	"testing"
)

//给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。 
//
// 输入为三个整数：day、month 和 year，分别表示日、月、年。 
//
// 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。 
//
// 
//
// 示例 1： 
//
// 输入：day = 31, month = 8, year = 2019
//输出："Saturday"
// 
//
// 示例 2： 
//
// 输入：day = 18, month = 7, year = 1999
//输出："Sunday"
// 
//
// 示例 3： 
//
// 输入：day = 15, month = 8, year = 1993
//输出："Sunday"
// 
//
// 
//
// 提示： 
//
// 
// 给出的日期一定是在 1971 到 2100 年之间的有效日期。 
// 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
func dayOfTheWeek(day int, month int, year int) string {
	weekDays := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// 别人的题意肯定不让用time包..
	if month < 3 {
		month += 12
		year -= 1
	}
	//基姆拉尔森计算公式
	//W= (d+2*m+3*(m+1)/5+y+y/4-y/100+y/400+1) mod 7
	//在公式中d表示日期中的日数，m表示月份数，y表示年数。
	//注意：在公式中有个与其他公式不同的地方：
	//把一月和二月看成是上一年的十三月和十四月，例：如果是2004-1-10则换算成：2003-13-10来代入公式计算。
	y := year
	d := day
	m := month
	w := (d+2*m+3*(m+1)/5+y+y/4-y/100+y/400+1)%7
	fmt.Println("W", w)
	return weekDays[w%7]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDayOfTheWeek(t *testing.T) {
	fmt.Println(dayOfTheWeek(18, 7, 1999) )
	fmt.Println(dayOfTheWeek(15, 8, 1993) )
	fmt.Println(dayOfTheWeek(30, 1, 2020) )
	fmt.Println(dayOfTheWeek(31, 8, 2020) )
}
