package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

var db *gorm.DB
var err error

func init() {
	dsn := "host=localhost user=aidos password=123 dbname=go port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	db.AutoMigrate(&User{}) // Автоматически создаст таблицу users
}
func insertUser(name string, age int) {
	user := User{Name: name, Age: age}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println("Unable to insert user:", result.Error)
	} else {
		log.Println("User inserted successfully:", user)
	}
}
func queryUsers() {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Println("Unable to query users:", result.Error)
		return
	}
	for _, user := range users {
		log.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
func main() {
	insertUser("Alice", 25)
	insertUser("Bob", 30)
	queryUsers()
}
