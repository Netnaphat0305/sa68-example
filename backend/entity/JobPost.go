package entity

import (
	"time"

	"gorm.io/gorm"
)

type StatusEnum string

const (
	Open  StatusEnum = "Open"
	Close StatusEnum = "Close"
)

type JobPost struct {
	gorm.Model
	Title       		string    		`gorm:"type:varchar(255);not null" json:"title"`
	Description 		string    		`gorm:"type:text;not null" json:"description"`
	Deadline    		time.Time 		`gorm:"type:date;not null" json:"deadline"`
	Status				StatusEnum 		`gorm:"not null" json:"status"`
	ImageURL 			*string 		`gorm:"type:varchar(255)" json:"image_url"` // nullable string
	PortfolioRequired 	*string 		`gorm:"type:varchar(100)" json:"portfolio_required"` // nullable
	Salary				int				`json:"salary"`

	EmployerID			uint    		`gorm:"not null" json:"employer_id"`
    Employer   			Employer 		`gorm:"foreignKey:EmployerID;references:ID"`

	JobCategoryID		uint			`gorm:"not null" json:"job_category_id"`
	JobCategory   		JobCategory 	`gorm:"foreignKey:JobCategoryID;references:ID"`

	LocationID			uint			`gorm:"not null" json:"location_id"`
	Location			Location		`gorm:"foreignKey:LocationID;references:ID"`

	EmploymentTypeID 	uint 			`gorm:"not null" json:"employment_type_id"`
	EmploymentType   	EmploymentType 	`gorm:"foreignKey:EmploymentTypeID;references:ID"`

	SalaryTypeID		uint			`gorm:"not null" json:"salary_type_id"`
	SalaryType			SalaryType		`gorm:"foreignKey:SalaryTypeID;references:ID"`

}
