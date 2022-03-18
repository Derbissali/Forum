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

func (handle *Dbs) PostPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/post_page.html", cfg.HTML+"header.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M internal.Forum

	M.Category = database.SelectCat(handle.Db)

	id := r.RequestURI[6:]
	// CommentId
	if !strings.HasPrefix(r.URL.Path, "/post/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	posts := database.CountPost(handle.Db)
	i, err := strconv.Atoi(r.URL.Path[6:])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if i < 1 || i > posts {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var e error
	M.Post.Rows, e = database.SinglePost(handle.Db, id)
	if e != nil {
		fmt.Println(e)
		return
	}

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		tmpl.Execute(w, M)
		return
	}

	M.User, err = database.ReadByUuid(handle.Db, c.Value)

	if M.User.Session == false {
		c, err := r.Cookie("session")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("wugalei")

		c.Name = "session"
		c.MaxAge = -1
		c.HttpOnly = true
		http.SetCookie(w, c)
	}

	tmpl.Execute(w, M)

}
