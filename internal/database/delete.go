package database

import (
	"database/sql"
	"fmt"
)

func DeleteCommentLike(c *sql.DB, n int, idComment string, idPost string) {
	_, err := c.Exec(`DELETE FROM comment_like_dislike WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteLike(c *sql.DB, idPost string, n int) {
	_, err := c.Exec(`DELETE FROM likeNdis where post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
