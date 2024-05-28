package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string  `json:"title" gorm:"size:200;index"`
	Content string  `json:"content"`
	Tags    *string `json:"tags"`
	UserId  int     `json:"userId"`
}
