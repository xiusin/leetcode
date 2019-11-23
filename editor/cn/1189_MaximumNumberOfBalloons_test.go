package cn

import (
	"fmt"
	"testing"
)

//给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
//
// 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
//
//
//
// 示例 1：
//
//
//
// 输入：text = "nlaebolko"
//输出：1
//
//
// 示例 2：
//
//
//
// 输入：text = "loonbalxballpoon"
//输出：2
//
//
// 示例 3：
//
// 输入：text = "leetcode"
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= text.length <= 10^4
// text 全部由小写英文字母组成
//
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func maxNumberOfBalloons(text string) int {
	var counter = map[byte]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for i := range text {
		if _, ok := counter[text[i]]; ok {
			counter[text[i]]++
		}
	}
	counter['l'] /= 2
	counter['o'] /= 2
	var minVal = counter['a']
	for _,v := range counter {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumNumberOfBalloons(t *testing.T) {
	//fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	//fmt.Println(maxNumberOfBalloons("leetcode"))
	//fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw"))
}
