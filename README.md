# 配置文件 #

## codeFileName ## 
`${question.frontendQuestionId}_$!velocityTool.camelCaseName(${question.titleSlug})_test`


## codeTemplate ##
```go
package cn

import "testing"
${question.content}
${question.code}

func Test$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T)  {
	
}
```