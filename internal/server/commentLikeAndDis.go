package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"strings"
	"text/template"
)

func (handle *Dbs) Comment_likedis(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("templates/post_page.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var M internal.Forum
	idPost := r.FormValue("postId")
	idComment := r.FormValue("comId")

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)
	n := M.User.ID
	l := r.FormValue("commnetLike")
	d := r.FormValue("commentDislike")
	if !strings.HasPrefix(r.URL.Path, "/commentLike/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if l != "" {
		a := database.CommentDislike(handle.Db, n, idComment, idPost)
		liked := database.CommentLike(handle.Db, n, idComment, idPost)

		fmt.Println(n, idPost, idComment)

		if liked != 0 {
			database.DeleteCommentLike(handle.Db, n, idComment, idPost)
		}
		if a != 0 {
			database.UpdateCommentLike(handle.Db, n, idComment, idPost)
		}
		if a == 0 && liked == 0 {
			database.InsertCommentLike(handle.Db, idPost, n, idComment)
		}

	} else if d != "" {
		b := database.CommentLike(handle.Db, n, idComment, idPost)
		disliked := database.CommentDislike(handle.Db, n, idComment, idPost)
		if disliked != 0 {
			database.DeleteCommentLike(handle.Db, n, idComment, idPost)

		}
		if b != 0 {
			database.UpdateCommentDisike(handle.Db, n, idComment, idPost)
		}
		if b == 0 && disliked == 0 {
			database.InsertCommentDisike(handle.Db, idPost, n, idComment)
		}
		
	}
	database.UpdateCommentLikeCount(handle.Db, idPost, idComment)
	database.UpdateCommentDislikeCount(handle.Db, idPost, idComment)
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
