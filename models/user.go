package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required,email" gorm:"unique"`
	Password string   `json:"password" binding:"required"`
	Role     UserRole `gorm:"type:varchar(100);not null"`
}

// TODO: exlude password, createdAt from getAllUsers

//func (u User) change(user User) {
//	if u.Role == ADMIN_ROLE{
//		user.Role = ADMIN_ROLE
//	}
//	return
//}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRole string

const (
	ADMIN_ROLE    UserRole = "admin"
	CUSTOMER_ROLE UserRole = "customer"
)
