# leetcode 目录 #
leetcode-cn刷题目录, 用于记录go语言解题, 里面有todo是做题时候的思路, 当然, 不会的, 抄过来的也有.

# 配置文件 #

> codeFileName
`${question.frontendQuestionId}_$!velocityTool.camelCaseName(${question.titleSlug})_test`

> codeTemplate
```go
package cn

import "testing"
${question.content}
${question.code}

// 单元测试函数
func Test$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T)  {
	
}
```