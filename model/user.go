package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id          int32  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Patronymic  string `json:"patronymic"`
	DateOfBirth string `json:"dateOfBirth"`
	About       string `json:"about"`
	Photo       string `json:"photo"`
	Company     string `json:"company"`
}

func (u *User) GetUserById(id int32, db *sql.DB) *User {

	err := db.QueryRow("SELECT id,firstname, lastname FROM users WHERE id=$1", id).Scan(u.Id, u.Firstname, u.Lastname)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}
