package database

import "database/sql"

func DbInit() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
