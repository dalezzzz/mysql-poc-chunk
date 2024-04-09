package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:example@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create chunk insert
	values := 70000
	chunkSize := 1000
	loops := values / chunkSize

	for i := 0; i < loops; i++ {
		users := make([]User, chunkSize)
		for j := 0; j < chunkSize; j++ {
			users[j] = User{Name: fmt.Sprintf("User %d", i*chunkSize+j)}
		}

		// Create multiple records
		if err := db.Create(&users).Error; err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("Insert completed")
}
