package main

import (
	"context"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pro struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ali-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()
	// Migrate the schema
	if err := db.AutoMigrate(&Pro{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	err = gorm.G[Pro](db).Create(ctx, &Pro{Code: "H3", Price: 323})
	product, err := gorm.G[Pro](db).Where("id = ? ", 2).First(ctx)
	_, err = gorm.G[Pro](db).Where("id = ? ", product.ID).Update(ctx, "price", 200)
	_, err = gorm.G[Pro](db).Where("id = ?", product.ID).Delete(ctx)
	_, err = gorm.G[Pro](db).Where("id > ?", 5).Delete(ctx)
}
