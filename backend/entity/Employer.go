package entity

import 
	"gorm.io/gorm"

type Employer struct {
	gorm.Model
	CompanyName   string    `gorm:"type:varchar(100);not null" json:"company_name"`
	ContactPerson string    `gorm:"type:varchar(100);not null" json:"contact_person"`
	Phone         string    `gorm:"type:varchar(20);not null" json:"phone"`
	Address       string    `gorm:"type:text;not null" json:"address"`

	UserID 	      uint 		`gorm:"not null" json:"user_id"`
	User          User 		`gorm:"foreignKey:UserID;references:ID" json:"user"`
}
