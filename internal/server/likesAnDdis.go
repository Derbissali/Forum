package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func (handle *Dbs) Likedis(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("templates/post_page.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M internal.Forum
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	idPost := r.RequestURI[10:]

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)

	n := M.User.ID
	l := r.FormValue("like")
	d := r.FormValue("dislike")
	if !strings.HasPrefix(r.URL.Path, "/likeNdis/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	i, err := strconv.Atoi(r.URL.Path[10:])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	posts := database.CountPost(handle.Db)
	if i < 1 || i > posts {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if l != "" {
		a := database.PostDislike(handle.Db, n, idPost)
		liked := database.PostLike(handle.Db, n, idPost)
		if liked != 0 {
			database.DeleteLike(handle.Db, idPost, n)
		}
		if a != 0 {
			database.UpdateLike(handle.Db, idPost, n)
		}
		if a == 0 && liked == 0 {
			database.InsertPostLike(handle.Db, idPost, n)
		}

	} else if d != "" {
		b := database.PostLike(handle.Db, n, idPost)
		disliked := database.PostDislike(handle.Db, n, idPost)
		if disliked != 0 {
			database.DeleteLike(handle.Db, idPost, n)
		}
		if b != 0 {
			database.UpdateDislike(handle.Db, idPost, n)
		}
		if b == 0 && disliked == 0 {
			database.InsertPostDislike(handle.Db, idPost, n)
		}
	}
	database.UpdateLikeCount(handle.Db, idPost)
	database.UpdateDislikeCount(handle.Db, idPost)
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
