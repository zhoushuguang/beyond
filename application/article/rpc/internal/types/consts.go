package types

const (
	SortPublishTime = iota
	SortLikeCount
)

const (
	DefaultPageSize = 20
	DefaultLimit    = 200

	DefaultSortLikeCursor = 1 << 30
)
