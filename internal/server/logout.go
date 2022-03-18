package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"text/template"
)

func (handle *Dbs) Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("gogosad")
	_, err := template.ParseFiles(cfg.HTML + "home_page.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/logout" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	M := internal.User{}

	c, err := r.Cookie("session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	//fmt.Printf("cookie: '%v'\n", c)

	c.Name = "session"
	c.MaxAge = -1
	c.HttpOnly = true

	//fmt.Printf("http.Value: '%v'\n", c)

	//fmt.Printf("c.Value: '%v'\n", c.Value)
	M, err = database.ReadByUuid(handle.Db, c.Value)

	fmt.Println(M, "-------------s")
	database.DeleteSession(handle.Db, M.ID)
	M.Session = false
	http.SetCookie(w, c)
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
