package models

import (
	"github.com/b-hurskyi/go-back-learn/db"
	"github.com/b-hurskyi/go-back-learn/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id;
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&u.ID)
	return err
}
