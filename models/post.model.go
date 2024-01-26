package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	UserID string
	Title  string
	Desc   string
	Image  string
	Likes  []string
}
