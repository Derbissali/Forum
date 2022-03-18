package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"strings"
	"text/template"
)

func (handle *Dbs) LikedPosts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.HTML+"liked_post.html", cfg.HTML+"header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/likedPosts" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var M internal.Forum

	M.Category = database.SelectCat(handle.Db)

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		fmt.Println("qwe6")
		http.Redirect(w, r, "/signin", 301)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)
	n := M.User.ID

	M.Post.Rows = database.ProfileLiked(handle.Db, n)

	tmpl.Execute(w, M)
}
func (handle *Dbs) CreatedPosts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/liked_post.html", cfg.HTML+"header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/createdPosts" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var M internal.Forum

	M.Category = database.SelectCat(handle.Db)

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "/signin", 301)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)
	n := M.User.ID

	M.Post.Rows = database.ProfileCreated(handle.Db, n)

	tmpl.Execute(w, M)
}
func (handle *Dbs) PostsByCat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.HTML+"postsByCat.html", cfg.HTML+"header.html")
	if !strings.HasPrefix(r.URL.Path, "/Category/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	var M internal.Forum

	M.Category = database.SelectCat(handle.Db)

	title := r.RequestURI[10:]
	//posts := database.ChooseCategory(handle.Db, title)
	var e error
	M.Post.Rows, e = database.Select(handle.Db, title)
	if e != nil {
		fmt.Println(e)
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
