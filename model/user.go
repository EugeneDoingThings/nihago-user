package model

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"log"
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
	query, args, err := sq.
		Select("id, firstname, lastname, patronymic, date_of_birth, about, photo, company_id").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		log.Fatalf("Unable to build SELECT query: %v", err)
		return nil
	}

	err = db.QueryRow(query, args...).Scan(
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
		log.Fatalf("Unable to execute SELECT query: %v", err)
		return nil
	}
	return u
}

func (u *User) GetUserList(db *sql.DB) []User {
	query, _, err := sq.
		Select("id, firstname, lastname, patronymic, date_of_birth, about, photo, company_id").
		From("users").
		ToSql()

	if err != nil {
		log.Fatalf("Unable to build SELECT query: %v", err)
		return nil
	}

	rows, _ := db.Query(query)

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
