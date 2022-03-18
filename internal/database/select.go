package database

import (
	"database/sql"
	"fmt"
	"forum/internal"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// ReadByUuid is ...
func ReadByUuid(c *sql.DB, uuid string) (internal.User, error) {
	row, err := c.Query(`SELECT user.id, user.name
	FROM user
	INNER JOIN session ON session.user_id=user.id
	WHERE uuid = ?`, uuid)
	var m internal.User
	if err != nil {
		fmt.Println(err)
		return m, nil
	}

	//	fmt.Println(uuid)
	n := internal.User{}
	for row.Next() {
		e := row.Scan(&n.ID, &n.Name)
		if e != nil {
			fmt.Println(err)
			return m, nil
		}

	}
	if n.ID != 0 {
		m.Session = true
	} else {
		m.Session = false

	}

	m.ID = n.ID
	m.Name = n.Name
	//	fmt.Println(m.ID, "from")
	return m, nil
}

// func (m *User) DeleteSession(){

// }
func Select(c *sql.DB, s string) ([]internal.Post, error) {
	// stmt, ok := Queries["Select#CatPosts"]
	// if !ok {
	// 	return errors.New("There is no query Select")
	// }
	rows, e := c.Query((`SELECT post.id, post.name, post.body, user.name, post.Image
	FROM category 
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN user ON user.id=post.user_id
	WHERE category.name = ?`), s)
	var m internal.Post
	if e != nil {
		return m.Rows, e
	}

	for rows.Next() {
		e = rows.Scan(&m.ID, &m.Name, &m.Body, &m.User.Name, &m.Image)
		if e != nil {
			return m.Rows, e
		}
		var a internal.Post
		a.ID = m.ID
		a.Name = m.Name
		a.Cat = CategoryByID(c, a.ID)

		a.Body = m.Body
		a.User.Name = m.User.Name
		a.Image = m.Image
		m.Rows = append(m.Rows, a)
	}
	return m.Rows, nil
}
func SinglePost(c *sql.DB, i string) ([]internal.Post, error) {
	row := c.QueryRow((`SELECT post.id, post.name, post.body, post.likes, post.dislikes, user.Name, post.Image
	FROM category 
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN user ON user.id=post.user_id
	WHERE post.id = ?`), i)
	var a internal.Post
	var m internal.Post
	e := row.Scan(&a.ID, &a.Name, &a.Body, &a.Likes, &a.Dislikes, &a.User.Name, &a.Image)
	if e != nil {
		return a.Rows, e
	}
	a.Cat = CategoryByID(c, a.ID)
	if e != nil {
		fmt.Println(e)
	}

	a.Comm = SelectComment(c, a.ID)
	if e != nil {
		fmt.Println(e)
	}

	m.Rows = append(m.Rows, a)

	return m.Rows, nil
}
func CategoryByID(c *sql.DB, postid int) []internal.Category {

	row, e := c.Query(`SELECT category.name  
	FROM category_post
	INNER JOIN category ON category.id = category_post.category_id
	WHERE post_id = ?`, postid)
	if e != nil {
		log.Println(e)
		return nil
	}
	var m internal.Post
	defer row.Close()
	for row.Next() {
		var a internal.Category
		e = row.Scan(&a.Name)
		if e != nil {
			log.Println(e)
			return nil
		}

		m.Cat = append(m.Cat, a)

	}
	return m.Cat
}

func All(c *sql.DB) ([]internal.Post, error) {
	rows, e := c.Query(`SELECT post.id, user.name, post.name, post.body, post.Image
	FROM post
	INNER JOIN user ON user.id=post.user_id`)
	var m internal.Post
	if e != nil {
		log.Println(e)
		return m.Rows, nil
	}

	for rows.Next() {
		var a internal.Post
		e = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if e != nil {
			log.Println(e)
			return a.Rows, nil
		}

		a.Cat = CategoryByID(c, a.ID)

		// fmt.Println(m.ID)

		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}

func SelectComment(c *sql.DB, i int) []internal.Comment {
	// stmt, ok := Queries["Select#CatPosts"]
	// if !ok {
	// 	return errors.New("There is no query Select")
	// }
	rows, e := c.Query((`SELECT comment.id, body, post_id, user.name, comment.likes, comment.dislikes
	FROM comment 
	INNER JOIN user ON user.id = comment.user_id
	WHERE post_id=?`), i)
	if e != nil {
		return nil
	}
	var m internal.Post
	for rows.Next() {
		var a internal.Comment
		e = rows.Scan(&a.ID, &a.Body, &a.Post.ID, &a.User.Name, &a.Likes, &a.Dislikes)
		if e != nil {
			log.Println(e)
			return nil
		}

		m.Comm = append(m.Comm, a)

	}

	return m.Comm
}

func ProfileLiked(c *sql.DB, n int) []internal.Post {
	rows, err := c.Query(`SELECT DISTINCT post.id, post.name, post.body, user.name, post.Image
	FROM category 
	INNER JOIN likeNdis ON post.id = likeNdis.post_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN user ON user.id = likeNdis.user_id WHERE likeNdis.like=1 AND user.id=?`, n)
	if err != nil {
		log.Println(err)
		return nil
	}
	var m internal.Post
	for rows.Next() {
		err = rows.Scan(&m.ID, &m.Name, &m.Body, &m.User.Name, &m.Image)
		if err != nil {
			log.Println(err)
			return nil
		}
		var a internal.Post
		a.ID = m.ID
		a.Name = m.Name
		a.Body = m.Body
		a.User.Name = m.User.Name
		a.Cat = CategoryByID(c, a.ID)
		a.Image = m.Image
		m.Rows = append(m.Rows, a)
	}

	return m.Rows
}
func ProfileCreated(c *sql.DB, n int) []internal.Post {
	rows, err := c.Query(`SELECT DISTINCT post.id, post.name, post.body, user.name, post.Image
	FROM category 
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN user ON user.id = post.user_id WHERE user.id=?`, n)
	if err != nil {
		log.Println(err)
		return nil
	}
	var m internal.Post
	for rows.Next() {
		err = rows.Scan(&m.ID, &m.Name, &m.Body, &m.User.Name, &m.Image)
		if err != nil {
			log.Println(err)
			return nil
		}
		var a internal.Post
		a.ID = m.ID
		a.Name = m.Name
		a.Body = m.Body
		a.User.Name = m.User.Name
		a.Cat = CategoryByID(c, a.ID)
		a.Image = m.Image
		m.Rows = append(m.Rows, a)
	}
	return m.Rows
}

func CommentDislike(c *sql.DB, n int, idComment string, idPost string) int {
	a := 0
	stmt := c.QueryRow(`SELECT "dislike" FROM "comment_like_dislike" WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	stmt.Scan(&a)
	return a
}
func CommentLike(c *sql.DB, n int, idComment string, idPost string) int {
	a := 0
	stmt := c.QueryRow(`SELECT "like" FROM "comment_like_dislike" WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	stmt.Scan(&a)
	return a
}
func SelectCat(c *sql.DB) []internal.Category {
	rows, e := c.Query(`SELECT "name" FROM "category" ORDER BY "name"`)
	if e != nil {
		return nil
	}
	var m internal.Category
	for rows.Next() {
		cat1 := internal.Category{}
		e = rows.Scan(&cat1.Name)
		if e != nil {
			return nil
		}
		m.Rows = append(m.Rows, cat1)

	}
	return m.Rows
}
func PostDislike(c *sql.DB, n int, idPost string) int {
	a := 0
	stmt := c.QueryRow(`SELECT "dislike" FROM "likeNdis" WHERE user_id=? AND post_id=?`, n, idPost)
	stmt.Scan(&a)
	return a
}

func PostLike(c *sql.DB, n int, idPost string) int {
	a := 0
	stmtl := c.QueryRow(`SELECT "like" FROM "likeNdis" WHERE user_id=? AND post_id=?`, n, idPost)
	stmtl.Scan(&a)
	return a
}

func CountPost(c *sql.DB) int {
	a := 0
	row := c.QueryRow(`SELECT COUNT(DISTINCT name) FROM post`)
	e := row.Scan(&a)
	if e != nil {
		return a
	}
	return a
}

func ChooseCategory(c *sql.DB, title string) string {
	posts := ""
	row := c.QueryRow(`SELECT name FROM category WHERE name=?`, title)
	e := row.Scan(&posts)
	if e != nil {
		fmt.Println(e)
		return posts
	}
	return posts
}
func SelectUserID(c *sql.DB, name string) int {
	a := 0
	row := c.QueryRow((`SELECT user.id FROM user WHERE user.name = ?`), name)
	e := row.Scan(&a)
	if e != nil {
		fmt.Println(e)
		return a
	}
	return a
}
