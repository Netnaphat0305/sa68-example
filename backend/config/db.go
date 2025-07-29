package config

import (
   "fmt"
   "time"
   "example.com/sa68-example/entity"
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
   return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")
	db = database
}

func SetupDatabase() {
	// Migrate Entities
	db.AutoMigrate(
		&entity.User{},
		&entity.Student{},
		&entity.ResumeFile{},
		&entity.Location{},
		&entity.JobPost{},
		&entity.JobCategory{},
		&entity.JobApplication{},
		&entity.EmploymentType{},
		&entity.Employer{},
	)

	// Seed EmploymentType
	employmentTypes := []entity.EmploymentType{
		{EmploymentTypeName: "FullTime"},
		{EmploymentTypeName: "PartTime"},
		{EmploymentTypeName: "Freelance"},
		{EmploymentTypeName: "Contract"},
	}
	for _, et := range employmentTypes {
		db.FirstOrCreate(&et, entity.EmploymentType{EmploymentTypeName: et.EmploymentTypeName})
	}

	// Seed JobCategory
	jobCategories := []entity.JobCategory{
		{CategoryName: "IT"},
		{CategoryName: "Marketing"},
		{CategoryName: "Finance"},
		{CategoryName: "Chef"},
	}
	for _, jc := range jobCategories {
		db.FirstOrCreate(&jc, entity.JobCategory{CategoryName: jc.CategoryName})
	}

	// Seed Location
	locations := []entity.Location{
		{LocationName: "SUT Campus", Address: "123 University Rd.", City: "Muang", Province: "Nakhon Ratchasima", ProvinceCode: 30000},
		{LocationName: "Bangkok Office", Address: "456 Bangkok St.", City: "Bangkok", Province: "Bangkok", ProvinceCode: 100000},
		{LocationName: "Lahuimala", Address: "456 University Rd.", City: "Muang", Province: "Nakhon Ratchasima", ProvinceCode: 30000},

	}
	for _, loc := range locations {
		db.FirstOrCreate(&loc, entity.Location{LocationName: loc.LocationName})
	}

	// Seed Users
	hashedPassword, _ := HashPassword("123456")

	userEmployer := entity.User{
		Username: "employer1",
		Password: hashedPassword,
		Role:     "employer",
	}
	db.FirstOrCreate(&userEmployer, entity.User{Username: userEmployer.Username})

	userStudent := entity.User{
		Username: "student1",
		Password: hashedPassword,
		Role:     "student",
	}
	db.FirstOrCreate(&userStudent, entity.User{Username: userStudent.Username})

	// Seed Employer
	employer := entity.Employer{
		CompanyName:   "ACME Corp",
		ContactPerson: "Mr. Boss",
		Phone:         "0812345678",
		Address:       "789 Business Rd.",
		UserID:        userEmployer.ID,
	}
	db.FirstOrCreate(&employer, entity.Employer{UserID: userEmployer.ID})

	// Seed Student
	birthDate, _ := time.Parse("2006-01-02", "2000-05-20")
	student := entity.Student{
		ID:        "B6608019",
		Email:     "student1@example.com",
		FirstName: "Netnaphat",
		LastName:  "Chamninok",
		Birthday:  birthDate,
		Gender:    "FeMale",
		Age:       20,
		GPA:       3.50,
		Year:      4,
		Faculty:   "Engineering",
		Phone:     "0899999999",
		UserID:    userStudent.ID,
	}
	db.FirstOrCreate(&student, entity.Student{ID: student.ID})

	// Seed JobPost
	jobPost := entity.JobPost{
		Title:            "Software Developer Intern",
		Description:      "Internship position for software development.",
		Deadline:         time.Now().AddDate(0, 1, 0),
		Status:           "Open",
		EmployerID:       employer.ID,
		JobCategoryID:    jobCategories[0].ID,
		LocationID:       locations[0].ID,
		EmploymentTypeID: employmentTypes[0].ID,
	}
	db.FirstOrCreate(&jobPost, entity.JobPost{Title: jobPost.Title})

	// Seed JobApplication
	jobApplication := entity.JobApplication{
		ApplicationStatus: "Pending",
		LastUpdate:        time.Now(),
		ApplicationReason: "I want to learn and contribute",
		StudentID:         student.ID,
		JobPostID:         jobPost.ID,
	}
	db.FirstOrCreate(&jobApplication, entity.JobApplication{StudentID: student.ID, JobPostID: jobPost.ID})

	// Seed ResumeFile
	resume := entity.ResumeFile{
		Filename:        "netnaphat_resume.pdf",
		UploadedDate:    time.Now(),
		JobApplicationID: jobApplication.ID,
	}
	db.FirstOrCreate(&resume, entity.ResumeFile{Filename: resume.Filename})
}
