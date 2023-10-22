package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type FollowCount struct {
	ID          int64 `gorm:"primary_key"`
	UserID      int64
	FollowCount int
	FansCount   int
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (m *FollowCount) TableName() string {
	return "follow_count"
}

type FollowCountModel struct {
	db *gorm.DB
}

func NewFollowCountModel(db *gorm.DB) *FollowCountModel {
	return &FollowCountModel{
		db: db,
	}
}

func (m *FollowCountModel) Insert(ctx context.Context, data *FollowCount) error {
	return m.db.Create(data).Error
}

func (m *FollowCountModel) FindOne(ctx context.Context, id int64) (*FollowCount, error) {
	var result FollowCount
	err := m.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *FollowCountModel) Update(ctx context.Context, data *FollowCount) error {
	return m.db.Save(data).Error
}

func (m *FollowCountModel) IncrFollowCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("INSERT INTO follow_count (user_id, follow_count) VALUES (?, 1) ON DUPLICATE KEY UPDATE follow_count = follow_count + 1", userId).
		Error
}

func (m *FollowCountModel) DecrFollowCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("UPDATE follow_count SET follow_count = follow_count - 1 WHERE user_id = ? AND follow_count > 0", userId).
		Error
}

func (m *FollowCountModel) IncrFansCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("INSERT INTO follow_count (user_id, fans_count) VALUES (?, 1) ON DUPLICATE KEY UPDATE fans_count = fans_count + 1", userId).
		Error
}

func (m *FollowCountModel) DecrFansCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("UPDATE follow_count SET fans_count = fans_count - 1 WHERE user_id = ? AND fans_count > 0", userId).
		Error
}

func (m *FollowCountModel) FindByUserIds(ctx context.Context, userIds []int64) ([]*FollowCount, error) {
	var result []*FollowCount
	err := m.db.WithContext(ctx).Where("user_id IN ?", userIds).Find(&result).Error
	return result, err
}
