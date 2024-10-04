package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Определение моделей User и Profile
type User struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `gorm:"unique;not null" json:"name"`
	Age     int     `json:"age"`
	Profile Profile `gorm:"constraint:OnDelete:CASCADE;" json:"profile"`
}

type Profile struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	UserID            uint   `gorm:"not null" json:"user_id"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
}

// Инициализация базы данных
func initDB() {
	dsn := "host=localhost user=aidos password=123 dbname=go port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Автоматическая миграция
	db.AutoMigrate(&User{}, &Profile{})
}

// Вставка пользователя с профилем
func createUserWithProfile(w http.ResponseWriter, r *http.Request) {
	var user User

	// Декодируем JSON в структуру
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Начинаем транзакцию
	tx := db.Begin()

	// Сохраняем пользователя
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Заполняем и сохраняем профиль
	profile := Profile{
		UserID:            user.ID,
		Bio:               user.Profile.Bio,
		ProfilePictureURL: user.Profile.ProfilePictureURL,
	}
	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	fmt.Fprintf(w, "User with profile created: %+v", user)
}

// Получение всех пользователей с их профилями
func getUsersWithProfiles(w http.ResponseWriter, r *http.Request) {
	var users []User

	if err := db.Preload("Profile").Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Обновление профиля пользователя
func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	params := mux.Vars(r)
	id := params["id"]

	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Находим профиль по UserID
	if err := db.Model(&Profile{}).Where("user_id = ?", id).Updates(profile).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Profile updated for user ID: %s", id)
}

// Удаление пользователя с профилем
func deleteUserWithProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	tx := db.Begin()

	// Удаляем профиль
	if err := tx.Where("user_id = ?", id).Delete(&Profile{}).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Удаляем пользователя
	if err := tx.Delete(&User{}, id).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()
	fmt.Fprintf(w, "User and profile deleted for user ID: %s", id)
}

func main() {
	initDB()

	router := mux.NewRouter()
	router.HandleFunc("/user", createUserWithProfile).Methods("POST")
	router.HandleFunc("/users", getUsersWithProfiles).Methods("GET")
	router.HandleFunc("/user/{id}/profile", updateUserProfile).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUserWithProfile).Methods("DELETE")

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Для создания пользователя с профилем:
// POST: http://localhost:8000/user
// Для получения всех пользователей с профилями:
// GET: http://localhost:8000/users
// Для обновления профиля пользователя (где {id} - это ID пользователя):
// PUT: http://localhost:8000/user/{id}/profile
// Для удаления пользователя с профилем (где {id} - это ID пользователя):
// DELETE: http://localhost:8000/user/{id}

// {
// 	"name": "Alice",
// 	"age": 25,
// 	"profile": {
// 	  "bio": "Software Developer",
// 	  "profile_picture_url": "http://example.com/image.jpg"
// 	}
//   }
