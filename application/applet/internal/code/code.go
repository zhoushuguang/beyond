package code

import "beyond/pkg/xcode"

var (
	RegisterMobileEmpty   = xcode.New(10001, "注册手机号不能为空")
	VerificationCodeEmpty = xcode.New(100002, "验证码不能为空")
	MobileHasRegistered   = xcode.New(100003, "手机号已经注册")
	LoginMobileEmpty      = xcode.New(100003, "手机号不能为空")
	RegisterPasswdEmpty   = xcode.New(100004, "密码不能为空")
)
