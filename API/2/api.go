package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type EnrollmentStatus string

var db *gorm.DB

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
	CanceledAt *time.Time       `gorm:"not null"`
	EnrolledAt *time.Time       `gorm:"not null"`
	StudentId  int
	CourseId   int
}

func Createuser(c fiber.Ctx) error {
	user := new(Students)

	if err := db.Create(user).Error; err != nil {
		// GORM error during creation (e.g., unique constraint violation)
		log.Printf("GORM Error creating user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database insertion failed",
		})
	}

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"massage": user,
		"id":      user.ID,
	})

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

	err = gorm.G[Students](db).Create(ctx, &Students{StudentCode: "4141414141", FirstName: "Mahdi", LastName: "Miladi"})
	err = gorm.G[Courses](db).Create(ctx, &Courses{CourseCode: "000", Title: "Arabic", Capacity: 123, EnrolledCount: 1, IsActive: true})
	app.Post("/api/v1/students", Createuser)

	log.Fatal(app.Listen(":3000"))
}
