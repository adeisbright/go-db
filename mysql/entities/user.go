package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `json:"username"`
	Email    string `json:"email" gorm:"size:50;index:idx_name,unique"`
}
