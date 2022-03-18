package database

import (
	"database/sql"
	"fmt"
)

func DeleteSession(c *sql.DB, n int) {
	_, err := c.Exec(`DELETE FROM session where user_id = ?`, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func AddSession(c *sql.DB, val string, n int) error {
	_, err := c.Exec(`INSERT INTO session (uuid, user_id) VALUES (?, ?)`, val, n)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SelectUser2(c *sql.DB, name string) int {
	a := 0
	row := c.QueryRow((`SELECT user_id FROM session 
	INNER JOIN user ON user.id=session.user_id
	WHERE user.name = ?`), name)
	e := row.Scan(&a)
	if e != nil {
		fmt.Println(e)
		return a
	}
	return a
}
