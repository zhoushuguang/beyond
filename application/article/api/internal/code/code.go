package code

import "beyond/pkg/xcode"

var (
	GetBucketErr              = xcode.New(30001, "获取bucket实例失败")
	PutBucketErr              = xcode.New(30002, "上传bucket失败")
	GetObjDetailErr           = xcode.New(30003, "获取对象详细信息失败")
	ArtitleTitleEmpty         = xcode.New(30004, "文章标题为空")
	ArticleContentTooFewWords = xcode.New(30005, "文章内容字数太少")
	ArticleCoverEmpty         = xcode.New(30006, "文章封面为空")
)
