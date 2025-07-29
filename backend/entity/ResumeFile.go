package entity

import (
	"time"
	"gorm.io/gorm"
)

type ResumeFile struct {
	gorm.Model
	Filename        string          `gorm:"type:varchar(100);not null" json:"filename"`
	UploadedDate    time.Time       `gorm:"not null" json:"uploaded_date"`
	
	JobApplicationID uint           `gorm:"not null" json:"job_application_id"`
	JobApplication   JobApplication `gorm:"foreignKey:JobApplicationID" json:"job_application"`

}
