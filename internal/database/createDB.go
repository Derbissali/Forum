package database

import (
	"database/sql"
	"fmt"
)

func Connect(db *sql.DB) error {

	CategoryDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "category" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	CategoryDb.Exec()
	category_postDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "category_post" (
		"id"	INTEGER NOT NULL,
		"category_id"	INTEGER NOT NULL,
		"post_id"	INTEGER NOT NULL,
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("category_id") REFERENCES "category"("id")
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	category_postDb.Exec()
	commentDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "comment" (
		"id"	INTEGER NOT NULL,
		"body"	TEXT NOT NULL,
		"post_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"likes"	INTEGER,
		"dislikes"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	commentDb.Exec()
	likeNdisDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "likeNdis" (
		"id"	INTEGER NOT NULL,
		"like"	INTEGER,
		"dislike"	INTEGER,
		"post_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		FOREIGN KEY("user_id") REFERENCES "user"("id"),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	likeNdisDb.Exec()
	postDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "post" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		"body"	TEXT NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"likes"	INTEGER,
		"dislikes"	INTEGER,
		"image"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	postDb.Exec()
	sessionDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "session" (
		"uuid"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL UNIQUE
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sessionDb.Exec()
	userDb, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "user" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		"email"	TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	userDb.Exec()
	comment_like_dislike, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "comment_like_dislike" (
		"id"	INTEGER NOT NULL,
		"like"	INTEGER,
		"dislike"	INTEGER,
		"comment_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"post_id"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("comment_id") REFERENCES "comment"("id"),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comment_like_dislike.Exec()
	stmt, err := db.Prepare(`
	INSERT INTO category(name) VALUES("anime"), ("marvel");`)
	if err != nil {
		return nil
	}
	stmt.Exec()

	return nil
}
