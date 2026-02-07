package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
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
}

type Courses struct {
	gorm.Model
	CourseCode    string `gorm:"size:32;unique;not null"`
	Title         string `gorm:"size:128;not null"`
	Capacity      int    `gorm:"not null"`
	EnrolledCount int    `gorm:"default:0;not null"`
	IsActive      bool   `gorm:"default:true;not null"`
}

type Enrollments struct {
	gorm.Model
	Status     EnrollmentStatus `gorm:"size : 10;default:'enrolled';not null"`
	CanceledAt *time.Time       `gorm:"null"`
	EnrolledAt *time.Time       `gorm:"not null"`
	StudentId  int
	CourseId   int
}

func main() {
	app := fiber.New()
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

	err = gorm.G[Students](db).Create(ctx, &Students{StudentCode: "404161494", FirstName: "Mahdi", LastName: "Miladi"})
	err = gorm.G[Courses](db).Create(ctx, &Courses{CourseCode: "11", Title: "Arabic", Capacity: 123, EnrolledCount: 1, IsActive: true})
	err = gorm.G[Enrollments](db).Create(ctx, &Enrollments{Status: StatusEnrolled, CanceledAt: nil, StudentId: 1, CourseId: 1})

	app.Post("/api/v1/students", func(c *fiber.Ctx) error {

	})
}
