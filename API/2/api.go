package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	ID          uint   `gorm:"primaryKey;autoIncrement;not null"`
}

type Courses struct {
	gorm.Model
	CourseCode    string `gorm:"size:32;unique;not null"`
	Title         string `gorm:"size:128;not null"`
	Capacity      int    `gorm:"not null"`
	EnrolledCount int    `gorm:"default:0;not null"`
	IsActive      bool   `gorm:"not null;default:true"`
	ID            uint   `gorm:"primaryKey;autoIncrement;not null"`
}

type Enrollments struct {
	gorm.Model
	Status     EnrollmentStatus `gorm:"size : 10;default:'enrolled';not null"`
	CanceledAt *time.Time       `gorm:"not null"`
	EnrolledAt *time.Time       `gorm:"not null"`
	StudentId  uint             `gorm:"unique;not null"`
	CourseId   uint             `gorm:"not null"`
	ID         uint             `gorm:"primaryKey;autoIncrement;not null"`
}

func database() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ali-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateUser(c fiber.Ctx) error {
	students := new(Students)

	db2 := database()

	if err := c.Bind().JSON(students); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db2.Create(students).Error; err != nil {
		fmt.Println("err", err, 12*22)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(students)

}

func listUsers(c fiber.Ctx) error {
	var students []Students

	db2 := database()

	if err := db2.Find(&students, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(students)
}

func GetUsers(c fiber.Ctx) error {
	var students Students

	db2 := database()
	if err := db2.Find(&students, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(students)
}

func CreateCourse(c fiber.Ctx) error {
	var course Courses

	db2 := database()

	if err := c.Bind().JSON(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db2.Create(&course).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(course)
}

func UpdateUser(c fiber.Ctx) error {

	updateData := Students{
		StudentCode: "40506070809",
		FirstName:   "John",
		LastName:    "Smith",
	}
	var students Students

	db2 := database()

	if err := db2.Model(&students).Where("id = ?", c.Params("id")).Updates(updateData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(students)
}

func UpdateUser2(c fiber.Ctx) error {
	var students Students

	db2 := database()

	d := c.Params("id")
	f, _ := strconv.Atoi(d)
	updateUser := Students{
		ID:          uint(f),
		StudentCode: "0928428941",
		FirstName:   "Ali",
		LastName:    "Ali",
	}

	if err := db2.Model(&Students{ID: updateUser.ID}).Updates(&updateUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(students)

}

func DeleteUser(c fiber.Ctx) error {
	var students Students

	db2 := database()

	if err := db2.Delete(&students, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(students)

}

func ListCourses(c fiber.Ctx) error {
	var courses []Courses

	db2 := database()

	if err := db2.Find(&courses, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error1": err.Error()})
	}
	return c.Status(200).JSON(courses)

}

func Recourses(c fiber.Ctx) error {
	var courses Courses

	db2 := database()

	if err := db2.Find(&courses, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error2": err.Error()})
	}
	return c.JSON(courses)
}

func DeleteCourse(c fiber.Ctx) error {
	var courses Courses

	db2 := database()

	if err := db2.Delete(&courses, c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(courses)
}

func UpdateCourse(c fiber.Ctx) error {
	couserData := map[string]interface{}{
		"is_active": false,
	}
	var courses Courses

	db2 := database()

	if err := db2.Model(&courses).Where("id = ?", c.Params("id")).Updates(couserData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(courses)

}

func UpdateCourse2(c fiber.Ctx) error {
	var courses Courses

	db2 := database()

	d := c.Params("id")
	f, _ := strconv.Atoi(d)
	asd := Courses{
		ID:         uint(f),
		CourseCode: c.Params("course_code"),
		Title:      c.Params("title"),
	}
	if err := db2.Model(&Courses{ID: asd.ID}).Updates(&asd).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(courses)
}

func CreateEnrollment(c fiber.Ctx, tx *gorm.DB) error {
	var enrollment Enrollments
	var courses Courses

	if err := c.Bind().JSON(&enrollment); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	student := Students{
		ID: enrollment.StudentId,
	}
	course := Courses{
		ID:            enrollment.CourseId,
		EnrolledCount: 0,
	}
	if err := tx.First(&student, enrollment.StudentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"code": 404, "message": "student not found"})
		}
	}
	if err := tx.First(&course, enrollment.CourseId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(409).JSON(fiber.Map{"code": 409, "message": "course not found"})
		}
	}

	if err := tx.Clauses(clause.Locking{Strength: "Update"}).First(&course, course.EnrolledCount).Error; err != nil {
		return c.Status(409).JSON(fiber.Map{"error": err.Error()})
	}

	if courses.Capacity <= (course).EnrolledCount {
		return c.Status(409).JSON(fiber.Map{"error": "capacity is completed"})
	}

	(course).EnrolledCount++

	if err := tx.Create(&enrollment).Error; err != nil {
		return c.Status(409).JSON(fiber.Map{"error": err.Error()})
	}
	if err := tx.Save(&courses).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})
	}

	return c.Status(200).JSON(enrollment)
}

func ErrorHandler(t fiber.Ctx) error {
	db2 := database()

	err := db2.Transaction(func(tx *gorm.DB) error {
		return CreateEnrollment(t, db2)
	})
	if err != nil {
		return t.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return t.Status(200).JSON(fiber.Map{"ok": true})
}

func main() {
	app := fiber.New()
	db := database()

	err := db.AutoMigrate(&Students{}, &Courses{}, &Enrollments{})
	if err != nil {
		panic("failed to connect database")
	}

	api := app.Group("/api")
	api.Post("/v1/students", CreateUser)
	api.Get("/v2/students", listUsers)
	api.Get("/v1/students/:id", GetUsers)
	api.Patch("/v1/students/:id", UpdateUser)
	api.Delete("/v1/students/:id", DeleteUser)
	api.Post("/v1/courses", CreateCourse)
	api.Get("/v1/courses", ListCourses)
	api.Get("/v1/courses/:id", Recourses)
	api.Delete("/v1/courses/:id", DeleteCourse)
	api.Patch("/v1/courses/:id", UpdateCourse)
	api.Put("/v1/students/:id", UpdateUser2)
	api.Put("/v1/courses/:id", UpdateCourse2)
	api.Post("/v1/enrollment", ErrorHandler)

	log.Fatal(app.Listen(":3000"))

}
