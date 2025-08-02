package config

import (
   "fmt" //สำหรับแสดงข้อความ เช่น fmt.Println
   "time"
   "example.com/sa68-example/entity"
   "gorm.io/driver/sqlite" //GORM ORM + SQLite driver
   "gorm.io/gorm"
)

var db *gorm.DB //สร้างตัวแปร db เพื่อเก็บ object ของ database 
//ฟังก์ชัน getter เพื่อให้ module อื่นเรียกใช้งาน database connection ได้ผ่าน config.DB()
func DB() *gorm.DB {
   return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		//ถ้าเชื่อมไม่สำเร็จ จะหยุดทำงานด้วย panic(...)
	}
	fmt.Println("Connected to database")
	//ถ้าเชื่อมได้จะพิมพ์ว่า "Connected to database" และกำหนดให้ตัวแปร db ชี้ไปยัง database ที่เชื่อมอยู่
	db = database
}

func SetupDatabase() {
	// Migrate Entities สร้างตารางตามโครงสร้าง struct ที่กำหนดไว้ใน entity
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
		&entity.SalaryType{},
	)

	// Seed EmploymentType สร้าง slice ของชนิดงาน เช่น งานประจำ (FullTime), พาร์ทไทม์ เหมือนเพิ่ทข้อมูล
	employmentTypes := []entity.EmploymentType{
		{EmploymentTypeName: "FullTime"},
		{EmploymentTypeName: "PartTime"},
		{EmploymentTypeName: "Freelance"},
		{EmploymentTypeName: "Contract"},
	}
	//FirstOrCreate จะเช็คก่อนว่ามีหรือยัง ถ้ายังไม่มีจะสร้างใหม่
	for _, et := range employmentTypes {
		db.FirstOrCreate(&et, entity.EmploymentType{EmploymentTypeName: et.EmploymentTypeName})
	}

	// Seed SalaryType
	salaryTypes := []entity.SalaryType{
		{SalaryTypeName: entity.Monthly},
		{SalaryTypeName: entity.Hourly},
		{SalaryTypeName: entity.Daily},
		{SalaryTypeName: entity.ProjectBased},
	}
	for _, st := range salaryTypes {
	db.FirstOrCreate(&st, entity.SalaryType{SalaryTypeName: st.SalaryTypeName})
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
	db.FirstOrCreate(&student, entity.Student{Email: student.Email})

	// Seed JobPost
		jobPost := entity.JobPost{
		Title:             "Software Developer Intern",
		Description:       "Internship position for software development.",
		Deadline:          time.Now().AddDate(0, 1, 0),
		Status:            "Open",
		ImageURL:          ptrString("https://yourdomain.com/uploads/job-image.png"),
		PortfolioRequired: ptrString("Required.pdf"),
		Salary:            12000,
		EmployerID:        employer.ID,
		JobCategoryID:     jobCategories[0].ID,
		LocationID:        locations[0].ID,
		EmploymentTypeID:  employmentTypes[0].ID,
		SalaryTypeID:      salaryTypes[0].ID,
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

	// ptrString แปลง string ให้เป็น pointer

}
func ptrString(s string) *string {
	return &s
}
