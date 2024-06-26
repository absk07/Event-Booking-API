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

func (e *Event) Save() error {
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

func GetEventById(id string) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = $1
	`
	row :=  db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5 
	`
	stmt, err := db.DB.Prepare(query)
	// fmt.Println("Query statement", stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}


func (e *Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = $1
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id)
	return err
}

func (e *Event) RegisterInEvent(userId string) error {
	query := `
		INSERT INTO registrations(id, event_id, user_Id) VALUES($1, $2, $3)
	`
	stmt, err := db.DB.Prepare(query)
	// fmt.Println("Query statement", stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.New().String(), e.Id, userId)
	return err
}

func (e *Event) CancleRegistration(userId string) error {
	query := `
		DELETE FROM registrations WHERE event_id = $1 AND user_id = $2
	`
	stmt, err := db.DB.Prepare(query)
	// fmt.Println("Query statement", stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, userId)
	return err
}