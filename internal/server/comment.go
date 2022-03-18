package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"text/template"
)

func (handle *Dbs) Commenting(w http.ResponseWriter, r *http.Request) {
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

	id := r.FormValue("idwka")

	// CommentId := r.FormValue("comIdd")
	// fmt.Println(CommentId)
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		fmt.Println("qwe2")
		http.Redirect(w, r, "/signin", 301)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)

	n := M.User.ID

	comment := r.FormValue("comment")
	if len(comment) > 0 && len(comment) < 140 {
		database.InsertComment(handle.Db, comment, id, n)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
