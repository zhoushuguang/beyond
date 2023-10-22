package code

import "beyond/pkg/xcode"

var (
	FollowUserIdEmpty   = xcode.New(40001, "关注用户id为空")
	FollowedUserIdEmpty = xcode.New(40002, "被关注用户id为空")
	CannotFollowSelf    = xcode.New(40003, "不能关注自己")
	UserIdEmpty         = xcode.New(40004, "用户id为空")
)
