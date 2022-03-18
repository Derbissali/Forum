package database

import (
	"database/sql"
	"fmt"
)

func InsertComment(c *sql.DB, comment string, id string, n int) {
	_, err := c.Exec(`INSERT INTO comment (body, post_id, user_id, likes, dislikes) VALUES (?, ?, ?,?,?)`, comment, id, n, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
func InsertCommentLike(c *sql.DB, idPost string, n int, idComment string) {
	_, err := c.Exec(`INSERT INTO comment_like_dislike (like, post_id, user_id, comment_id) VALUES (?, ?, ?, ?)`, 1, idPost, n, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func InsertCommentDisike(c *sql.DB, idPost string, n int, idComment string) {
	_, err := c.Exec(`INSERT INTO comment_like_dislike (dislike, post_id, user_id, comment_id) VALUES (?, ?, ?, ?)`, 1, idPost, n, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func InsertPostLike(c *sql.DB, idPost string, n int) {
	_, err := c.Exec(`INSERT INTO likeNdis (like, post_id, user_id) VALUES (?, ?, ?)`, 1, idPost, n)
	if err != nil {
		fmt.Println(err)
	}
}
func InsertPostDislike(c *sql.DB, idPost string, n int) {
	_, err := c.Exec(`INSERT INTO likeNdis (dislike, post_id, user_id) VALUES (?, ?, ?)`, 1, idPost, n)
	if err != nil {
		fmt.Println(err)
	}
}
func InsertUser(c *sql.DB, name string, email string, pass string) error {
	_, err := c.Exec(`INSERT INTO user (name, email, password) VALUES (?, ?, ?)`, name, email, pass)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
