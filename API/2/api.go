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
	StatusCanceled EnrollmentStatus = "canceled"
	StatusEnrolled EnrollmentStatus = "enrolled"
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
	Status     EnrollmentStatus `gorm:"size : 10;not null"`
	CanceledAt time.Time
	EnrolledAt time.Time `gorm:"autoCreateTime:milli"`
	StudentId  uint      `gorm:"unique;not null"`
	CourseId   uint      `gorm:"not null"`
	ID         uint      `gorm:"primaryKey;autoIncrement;not null"`
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
	var course Courses
	var student Students

	if err := c.Bind().JSON(&enrollment); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := tx.First(&student, enrollment.StudentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"code": 404, "message": "student not found"})
		}
	}

	if err := tx.Clauses(clause.Locking{Strength: "Update"}).First(&course, "id = ?", enrollment.CourseId).Error; err != nil {
		return c.Status(409).JSON(fiber.Map{"error": err.Error()})
	}

	if course.Capacity <= (course).EnrolledCount {
		return c.Status(409).JSON(fiber.Map{"code": 409, "massage": "capacity is completed"})
	}

	enrollment.Status = StatusEnrolled

	if err := tx.Create(&enrollment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	course.EnrolledCount++

	if err := tx.Save(&course).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": err.Error()})
	}

	return c.Status(200).JSON(enrollment)
}

func Cancle(c fiber.Ctx, tx *gorm.DB) error {
	var enrollment Enrollments
	var course Courses

	if err := tx.First(&enrollment, "id = ?", c.Params("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"code": 404, "message": "student not found in enrollments"})
		}
	}
	if err := tx.First(&enrollment, enrollment.Status).Error; err != nil {
		if enrollment.Status == StatusCanceled {
			return c.Status(409).JSON(fiber.Map{"code": 409, "message": "enrollment is canceled"})
		}
	}

	f := c.Params("id")
	d, err := strconv.Atoi(f)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	cor := Enrollments{
		ID:        uint(d),
		StudentId: enrollment.StudentId,
		CourseId:  enrollment.CourseId,
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&course, cor.CourseId).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error() + "Bye"})
	}

	if course.Capacity < 0 {
		return c.Status(409).JSON(fiber.Map{"code": 409, "message": "student capacity can not be less than 0"})
	}

	enrollment.Status = StatusCanceled
	enrollment.CanceledAt = time.Now()

	if err := tx.Save(&enrollment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	course.EnrolledCount--

	if err := tx.Save(&course).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(enrollment)
}

func ErrorHandler(c fiber.Ctx) error {
	db2 := database()

	err := db2.Transaction(func(tx *gorm.DB) error {
		return CreateEnrollment(c, db2)
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return nil
}

func handlercacle(c fiber.Ctx) error {
	db2 := database()
	err := db2.Transaction(func(tx *gorm.DB) error {
		return Cancle(c, db2)
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return nil
}

func listEnrollments(c fiber.Ctx) error {
	db2 := database()
	var enrollments []Enrollments
	if err := db2.Find(&enrollments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(enrollments)
}

func main() {
	app := fiber.New()

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
	api.Post("v1/enrollment/:id/cancel", handlercacle)
	api.Get("/v1/enrollment", listEnrollments)

	log.Fatal(app.Listen(":3000"))

}
