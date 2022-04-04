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
	CompanyId   int32  `json:"company_id"`
}

func (u *User) GetUserById(id int32, db *sql.DB) *User {
	err := db.QueryRow(
		"SELECT id, firstname, lastname, patronymic, date_of_birth, about, photo, company_id "+
			"FROM users "+
			"WHERE id=$1", id).Scan(
		&u.Id,
		&u.Firstname,
		&u.Lastname,
		&u.Patronymic,
		&u.DateOfBirth,
		&u.About,
		&u.Photo,
		&u.CompanyId,
	)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}

func (u *User) GetUserList(db *sql.DB) []User {
	rows, _ := db.Query(
		"SELECT id, firstname, lastname, patronymic, date_of_birth, about, photo, company_id " +
			"FROM users")

	var users []User

	for rows.Next() {
		if err := rows.Scan(&u.Id,
			&u.Firstname,
			&u.Lastname,
			&u.Patronymic,
			&u.DateOfBirth,
			&u.About,
			&u.Photo,
			&u.CompanyId); err != nil {
			return nil
		}
		users = append(users, *u)
	}

	return users
}
