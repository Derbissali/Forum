package main

import (
	"fmt"
	"forum/internal/database"
	"forum/internal/server"
	"log"
	"net/http"
)

func main() {

	db, err := database.DbInit()
	if err != nil {
		log.Println(err)
		return
	}
	database.Connect(db)
	handle := server.CreateHandle(db)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", handle.Home_page)
	http.HandleFunc("/signup", handle.Signup)
	http.HandleFunc("/signin", handle.Signin)
	http.HandleFunc("/logout", handle.Logout)
	http.HandleFunc("/createpost", handle.AddPost)
	http.HandleFunc("/post/", handle.PostPage)
	http.HandleFunc("/likeNdis/", handle.Likedis)
	http.HandleFunc("/posting", handle.Posting)
	http.HandleFunc("/commenting", handle.Commenting)
	http.HandleFunc("/createdPosts", handle.CreatedPosts)
	http.HandleFunc("/likedPosts", handle.LikedPosts)
	http.HandleFunc("/Category/", handle.PostsByCat)
	http.HandleFunc("/commentLike/", handle.Comment_likedis)
	fmt.Println("Project is running in localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("listenServe error: ", err.Error())
	}
}
