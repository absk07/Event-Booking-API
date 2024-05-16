package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

var DB *sql.DB

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		panic("Error connecting to PostgreSQL DB")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	fmt.Println("Connected to PostgreSQL DB!")

	CreateTable()
}

func CreateTable() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(60) PRIMARY KEY,
			email VARCHAR(60) NOT NULL UNIQUE,
			password VARCHAR(60) NOT NULL
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Error creating Users Table")
	}
	fmt.Println("Events Table created!")
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(60) PRIMARY KEY,
			name VARCHAR(60) NOT NULL,
			description VARCHAR(60) NOT NULL,
			location VARCHAR(60) NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			userId VARCHAR(60) NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Error creating Events Table")
	}
	fmt.Println("Events Table created!")
}
