package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbSql *sql.DB
var dbGorm *gorm.DB
var err error

// Модель пользователя
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Подключение через database/sql
	connStr := "user=aidos password=123 dbname=go sslmode=disable"
	dbSql, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database with sql:", err)
	}

	// Подключение через GORM
	dsn := "host=localhost user=aidos password=123 dbname=go port=5432 sslmode=disable"
	dbGorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database with GORM:", err)
	}

	// Авто-миграция для GORM
	dbGorm.AutoMigrate(&User{})

	// Маршрутизатор Gorilla Mux
	router := mux.NewRouter()

	// Маршруты для database/sql
	router.HandleFunc("/sql/users", GetUsersSql).Methods("GET")
	router.HandleFunc("/sql/user", CreateUserSql).Methods("POST")

	// Маршруты для GORM
	router.HandleFunc("/gorm/users", GetUsersGorm).Methods("GET")
	router.HandleFunc("/gorm/user", CreateUserGorm).Methods("POST")

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Функции для работы через database/sql

// Получение пользователей через database/sql
func GetUsersSql(w http.ResponseWriter, r *http.Request) {
	rows, err := dbSql.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Создание пользователя через database/sql
func CreateUserSql(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := dbSql.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		http.Error(w, "Unable to insert user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Функции для работы через GORM

// Получение пользователей через GORM
func GetUsersGorm(w http.ResponseWriter, r *http.Request) {
	var users []User
	dbGorm.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Создание пользователя через GORM
func CreateUserGorm(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	dbGorm.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// http://localhost:8000/sql/users
