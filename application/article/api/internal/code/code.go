package code

import "beyond/pkg/xcode"

var (
	GetBucketErr    = xcode.New(30001, "获取bucket实例失败")
	PutBucketErr    = xcode.New(30002, "上传bucket失败")
	GetObjDetailErr = xcode.New(30003, "获取对象详细信息失败")
)
