package code

import (
	"beyond/pkg/xcode"
)

var (
	SortTypeInvalid = xcode.New(60001, "排序类型无效") // 排序类型无效
	UserIdInvalid   = xcode.New(60002, "用户ID无效") // 用户ID无效
)
