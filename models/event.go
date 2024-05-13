package models

import (
	"database/sql"
	"time"

	"example.com/event-booking-api/db"
	"github.com/google/uuid"
)

type Event struct {
	Id          string
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      string
}

func (e Event) Save() error {
	var err error
	var stmt *sql.Stmt
	query := `
		INSERT INTO events(id, name, description, location, dateTime, userId) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	stmt, err = db.DB.Prepare(query)
	// fmt.Println("Query statement", stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.New().String(), e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err = rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, err
}
