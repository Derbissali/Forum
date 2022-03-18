package server

import "database/sql"

type Dbs struct {
	Db *sql.DB
}

func CreateHandle(db *sql.DB) *Dbs {
	return &Dbs{
		Db: db,
	}
}
