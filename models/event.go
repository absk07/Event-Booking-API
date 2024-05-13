package models

import "time"

type Event struct {
	Id          string
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      string
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
