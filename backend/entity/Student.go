package entity

import (
	"time"
	"gorm.io/gorm"
)

type GenderEnum string

const (
	Male   GenderEnum = "Male"
	Female GenderEnum = "Female"
	Other  GenderEnum = "Other"
)

type Student struct {
	gorm.Model
	// ID 		  string     `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Email     string     `gorm:"type:varchar(255);not null" json:"email"`
	FirstName string     `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string     `gorm:"type:varchar(100);not null" json:"last_name"`
	Birthday  time.Time  `gorm:"not null" json:"birthday"`
	Gender    GenderEnum `gorm:"not null" json:"gender"`
	Age       int        `gorm:"not null" json:"age"`
	GPA       float32    `gorm:"not null" json:"gpa"`
	Year      int        `gorm:"not null" json:"year"`
	Faculty   string     `gorm:"type:varchar(255);not null" json:"faculty"`
	Phone     string     `gorm:"type:varchar(20);not null" json:"phone"`
	// FK
	UserID uint  `gorm:"not null" json:"user_id"`
	User   User  `gorm:"foreignKey:UserID" json:"user"`
}

