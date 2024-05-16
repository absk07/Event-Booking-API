package models

import (
	"errors"

	"example.com/event-booking-api/db"
	"example.com/event-booking-api/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(id, email, password) VALUES ($1, $2, $3)`
	stmt, err := db.DB.Prepare(query);
	if err != nil {
		return err
	}
	defer stmt.Close()
	var hashedPassword string
	hashedPassword, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid.New().String(), u.Email, hashedPassword)
	return err
}

func (u User) ValidateUser() error {
	query := `SELECT password FROM users WHERE email = $1`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("Invalid Credentials")
	}
	isPasswordValid := utils.IsPasswordValid(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("Invalid Credentials")
	}
	return nil
}