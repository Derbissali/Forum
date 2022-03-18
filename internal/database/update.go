package database

import (
	"database/sql"
	"fmt"
)

func UpdateCommentLike(c *sql.DB, n int, idComment string, idPost string) {
	_, err := c.Exec(`UPDATE comment_like_dislike SET like=1, dislike=NULL WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func UpdateCommentDisike(c *sql.DB, n int, idComment string, idPost string) {
	_, err := c.Exec(`UPDATE comment_like_dislike SET like=NULL, dislike=1 WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateLike(c *sql.DB, idPost string, n int) {
	_, err := c.Exec(`UPDATE likeNdis SET like=1, dislike=NULL WHERE post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func UpdateDislike(c *sql.DB, idPost string, n int) {
	_, err := c.Exec(`UPDATE likeNdis SET like=NULL, dislike=1 WHERE post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func UpdateLikeCount(c *sql.DB, id string) {
	a := 0
	row := c.QueryRow(`SELECT COUNT(like) FROM likeNdis WHERE post_id=?`, id)
	e := row.Scan(&a)
	if e != nil {
		return
	}
	_, err := c.Exec(`UPDATE post SET likes=? WHERE id=?`, a, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func UpdateDislikeCount(c *sql.DB, id string) {
	a := 0
	row := c.QueryRow(`SELECT COUNT(dislike) FROM likeNdis WHERE post_id=?`, id)
	e := row.Scan(&a)
	if e != nil {
		return
	}
	_, err := c.Exec(`UPDATE post SET dislikes=? WHERE id=?`, a, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateCommentLikeCount(c *sql.DB, id string, cid string) {
	a := 0
	Comrow := c.QueryRow(`SELECT COUNT(like) FROM comment_like_dislike WHERE post_id=? AND comment_id=?`, id, cid)
	e := Comrow.Scan(&a)
	if e != nil {
		return
	}

	_, err := c.Exec(`UPDATE comment SET likes=? WHERE id=?`, a, cid)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func UpdateCommentDislikeCount(c *sql.DB, id string, cid string) {
	a := 0
	Comrow := c.QueryRow(`SELECT COUNT(dislike) FROM comment_like_dislike WHERE post_id=? AND comment_id=?`, id, cid)
	e := Comrow.Scan(&a)
	if e != nil {
		return
	}

	_, err := c.Exec(`UPDATE comment SET dislikes=? WHERE id=?`, a, cid)
	if err != nil {
		fmt.Println(err)
		return
	}
}
