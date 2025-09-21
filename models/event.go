package models

import (
	"context"
	"time"

	"github.com/b-hurskyi/go-back-learn/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save(ctx context.Context) error {
	query := `
		INSERT INTO events(name, description, location, date_time, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	return db.DB.QueryRowContext(ctx, query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
}

func GetEventById(ctx context.Context, id int64) (*Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = $1`
	row := db.DB.QueryRowContext(ctx, query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
