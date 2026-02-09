package main

import (
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

func CreateUser(c fiber.Ctx) error {
	user := new(Students)

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db.Create(user).Error; err != nil {
		// GORM error during creation (e.g., unique constraint violation)
		log.Printf("GORM Error creating user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database insertion failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"massage": "User registered successfully",
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

	err = db.AutoMigrate(&Students{}, &Courses{}, &Enrollments{})
	if err != nil {
		panic("failed to connect database")
	}
	api := app.Group("/api")
	api.Post("/v1/students", CreateUser)

	log.Fatal(app.Listen(":3000"))
}
