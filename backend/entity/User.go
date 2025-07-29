package entity

import (
	"gorm.io/gorm"
)

// ENUM: กำหนดบทบาท
type RoleEnum string

const (
	Stu  RoleEnum = "student"
	Emp RoleEnum = "employer"
)

// User: ใช้ได้ทั้งนักศึกษาและผู้ว่าจ้าง
type User struct {
	gorm.Model
	Username string   `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password string   `gorm:"type:varchar(100);not null" json:"password"`
	Role     RoleEnum `gorm:"type:varchar(20);not null" json:"role"`
}
