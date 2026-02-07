package main

import (
	"context"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type EnrollmentStatus string

const (
	StatusEnrolled EnrollmentStatus = "enrolled"
	StatusCanceled EnrollmentStatus = "canceled"
)

type Students struct {
	gorm.Model
	StudentCode string `gorm:"size:32;unique;not null"`
	FirstName   string `gorm:"size:64;not null"`
	LastName    string `gorm:"size:64;not null"`
	Enrollment2 []Enrollments
}

type Courses struct {
	gorm.Model
	CourseCode    string `gorm:"size:32;unique;not null"`
	Title         string `gorm:"size:128;not null"`
	Capacity      int    `gorm:"not null"`
	EnrolledCount int    `gorm:"default:0;not null"`
	IsActive      bool   `gorm:"default:true;not null"`
	Enrollment2   []Enrollments
}

type Enrollments struct {
	gorm.Model
	Status     EnrollmentStatus `gorm:"size : 10;default:'enrolled';not null"`
	CanceledAt *time.Time       `gorm:"not null"`
	EnrolledAt *time.Time       `gorm:"not null"`
	StudentId  int
	CourseId   int
	Student    Students `gorm:"foreignkey:StudentId;not null"`
	Course     Courses  `gorm:"foreignkey:CourseId;not null"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ali-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	err = db.AutoMigrate(&Students{}, &Courses{}, &Enrollments{})
	if err != nil {
		panic("failed to connect database")
	}

	err = gorm.G[Students](db).Create(ctx, &Students{StudentCode: "404161432", FirstName: "Mahdi", LastName: "Miladi"})
	err = gorm.G[Courses](db).Create(ctx, &Courses{CourseCode: "180", Title: "Arabic", Capacity: 123, EnrolledCount: 1, IsActive: true})

}
