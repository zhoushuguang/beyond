package types

// CanalLikeMsg canal解析like binlog消息.
type CanalLikeMsg struct {
	Data []struct {
		ID         string `json:"id"`
		BizID      string `json:"biz_id"`
		ObjID      string `json:"obj_id"`
		LikeNum    string `json:"like_num"`
		DislikeNum string `json:"dislike_num"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
	} `json:"data"`
}

type CanalArticleMsg struct {
	Data []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Content     string `json:"content"`
		Description string `json:"description"`
		AuthorId    string `json:"author_id"`
		Status      string `json:"status"`
		CommentNum  string `json:"comment_num"`
		LikeNum     string `json:"like_num"`
		CollectNum  string `json:"collect_num"`
		ViewNum     string `json:"view_num"`
		ShareNum    string `json:"share_num"`
		TagIds      string `json:"tag_ids"`
		PublishTime string `json:"publish_time"`
		CreateTime  string `json:"create_time"`
		UpdateTime  string `json:"update_time"`
	}
}

type ArticleEsMsg struct {
	ArticleId   int64   `json:"article_id"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Description string  `json:"description"`
	AuthorId    int64   `json:"author_id"`
	AuthorName  string  `json:"author_name"`
	Status      int     `json:"status"`
	CommentNum  int64   `json:"comment_num"`
	LikeNum     int64   `json:"like_num"`
	CollectNum  int64   `json:"collect_num"`
	ViewNum     int64   `json:"view_num"`
	ShareNum    int64   `json:"share_num"`
	TagIds      []int64 `json:"tag_ids"`
	PublishTime string  `json:"publish_time"`
	CreateTime  string  `json:"create_time"`
	UpdateTime  string  `json:"update_time"`
}
