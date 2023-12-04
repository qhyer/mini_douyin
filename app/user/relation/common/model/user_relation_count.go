package model

import "gorm.io/gorm"

type UserRelationCount struct {
	gorm.Model
	UserId        int64
	FollowCount   int64
	FollowerCount int64
}
