package model

type UserRelationCount struct {
	UserId        int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	FollowCount   int64 `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64 `gorm:"column:follower_count" json:"follower_count"`
}
