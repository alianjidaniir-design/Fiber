package main

import (
	"context"
	"math/big"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type students struct {
	gorm.Model
	studentCode string `gorm:"default"`
	firstName   string
	lastName    string
}

type courses struct {
	gorm.Model
	courseCode    string
	title         string
	capacity      int
	enrolledCount int
	isActive      bool
}

type enrollments struct {
	gorm.Model
	studentId  big.Int
	courseId   big.Int
	status     string
	canceled   time.Time
	enrolledAt time.Time
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ali-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	err = db.AutoMigrate(&students{}, &courses{}, &enrollments{})
	if err != nil {
		panic("failed to connect database")
	}

	err = gorm.G[students](db).Create(ctx, &students{studentCode: "404151521", firstName: "Ali", lastName: "Anjidani"})

}
