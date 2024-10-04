package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "aidos"
	password = "123"
	dbname   = "go"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	createTable(db)             // Создаем таблицу
	insertUser(db, "Alise", 26) // Вставляем данные
	insertUser(db, "Bobe", 31)
	queryUsers(db) // Выполняем запросы и выводим результаты
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Unable to create table: %v", err)
	}
	fmt.Println("Table created successfully!")
}

func insertUser(db *sql.DB, name string, age int) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatalf("Unable to insert user: %v", err)
	}
	fmt.Println("User inserted successfully!")
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatalf("Unable to query users: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users in database:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
