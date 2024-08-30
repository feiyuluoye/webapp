package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
