package entity

import (
	"time"
	"gorm.io/gorm"
)

type ApplicationStatusEnum string

const (
	StatusPending       ApplicationStatusEnum = "Pending"         // รอพิจารณา
	StatusInterview  	ApplicationStatusEnum = "InterviewPending"   //รอสัม
	StatusAccepted      ApplicationStatusEnum = "Accepted"        // ผ่านการคัดเลือก
    StatusRejected      ApplicationStatusEnum = "Rejected"        // ไม่ผ่านการคัดเลือก
)

type JobApplication struct {
    gorm.Model

    ApplicationStatus  ApplicationStatusEnum `gorm:"type:varchar(50);not null" json:"application_status"`
    LastUpdate        time.Time              `gorm:"not null" json:"last_update"`
    ApplicationReason string                 `gorm:"type:varchar(255);not null" json:"application_reason"`

    StudentID uint      `gorm:"not null" json:"student_id"`
    Student   Student `gorm:"foreignKey:StudentID;references:ID"`  // FK to Student.ID

    JobPostID uint    `gorm:"not null" json:"job_post_id"`
    JobPost   JobPost `gorm:"foreignKey:JobPostID;references:ID"`  // FK to JobPost.ID
}
