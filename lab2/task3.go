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

var gormDB *gorm.DB
var sqlDB *sql.DB

// Определение модели User
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	Age  int    `json:"age" gorm:"not null"`
}

// Инициализация базы данных
func initDB() {
	dsn := "host=localhost user=aidos password=123 dbname=go port=5432 sslmode=disable"
	var err error

	// Подключение к GORM
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to GORM database!", err)
	}

	// Подключение к SQL DB
	sqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to SQL database!", err)
	}

	// Проверка подключения
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping SQL database!", err)
	}

	// Автоматическая миграция
	err = gormDB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// Получение пользователей с фильтрацией и сортировкой (прямой SQL)
func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	age := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")

	var query string
	if age != "" {
		query = fmt.Sprintf("SELECT * FROM users WHERE age = %s", age)
	} else {
		query = "SELECT * FROM users"
	}

	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortBy)
	}

	rows, err := sqlDB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Создание пользователя (прямой SQL)
func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := sqlDB.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Обновление пользователя (прямой SQL)
func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := sqlDB.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", user.Name, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Удаление пользователя (прямой SQL)
func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := sqlDB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Получение пользователей с фильтрацией и сортировкой (GORM)
func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	age := r.URL.Query().Get("age")
	sortBy := r.URL.Query().Get("sort")

	query := gormDB.Model(&User{})
	if age != "" {
		query = query.Where("age = ?", age)
	}
	if sortBy != "" {
		query = query.Order(sortBy)
	}

	if err := query.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Создание пользователя (GORM)
func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := gormDB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Обновление пользователя (GORM)
func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := gormDB.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Удаление пользователя (GORM)
func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := gormDB.Delete(&User{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()

	router := mux.NewRouter()

	// Роуты для прямых SQL-запросов
	router.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
	router.HandleFunc("/sql/users", createUserSQL).Methods("POST")
	router.HandleFunc("/sql/users/{id}", updateUserSQL).Methods("PUT")
	router.HandleFunc("/sql/users/{id}", deleteUserSQL).Methods("DELETE")

	// Роуты для GORM
	router.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	router.HandleFunc("/gorm/users", createUserGORM).Methods("POST")
	router.HandleFunc("/gorm/users/{id}", updateUserGORM).Methods("PUT")
	router.HandleFunc("/gorm/users/{id}", deleteUserGORM).Methods("DELETE")

	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// http://localhost:8000/sql/users
// {
// 	"name": "John Doe",
// 	"age": 30
//   }
