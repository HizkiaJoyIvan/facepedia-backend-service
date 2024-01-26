package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string
	Password       string
	Email          string
	ProfilePicture string
	CoverPicture   string
	Followers      []string
	Followings     []string
	Desc           string
	City           string
}
