package database

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgreSQL struct {
	Instance *sql.DB
}

func (p *PostgreSQL) Init(user, password, dbname, port string) *sql.DB {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)

	var err error
	p.Instance, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return p.Instance
}
