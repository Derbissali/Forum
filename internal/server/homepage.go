package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"net/http"
	"text/template"
)

func (handle *Dbs) Home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.HTML+"home_page.html", cfg.HTML+"header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// if r.Method != http.MethodGet {
	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
	M := internal.Forum{}

	M.Category = database.SelectCat(handle.Db)
	var e error
	M.Post.Rows, e = database.All(handle.Db)

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

	// fmt.Println(M.Post)

	tmpl.Execute(w, M)
}
