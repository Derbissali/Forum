package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func (handle *Dbs) Signup(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.HTML + "signup.html")
	if err != nil {
		fmt.Println(err)
	}
	if r.URL.Path != "/signup" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		return
	}
	_, err = r.Cookie("session")
	if err == nil {
		log.Println("signin")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	creds := internal.User{}
	switch r.Method {
	case "GET":
		tmpl.Execute(w, creds)
	case "POST":
		creds := &internal.User{
			Name:     r.FormValue("login"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
		if handle.CheckEmpty(creds) != true {
			fmt.Println("empty")
			w.WriteHeader(http.StatusBadRequest)
			tmpl.Execute(w, creds)
			return
		}

		if handle.CheckUniq(creds) != true {
			fmt.Println("not uniq")
			w.WriteHeader(http.StatusBadRequest)
			tmpl.Execute(w, creds)
			return
		}
		if handle.CheckEmail(creds) != true {
			fmt.Println("email shit")
			w.WriteHeader(http.StatusBadRequest)
			tmpl.Execute(w, creds)
			return
		}
		err = database.InsertUser(handle.Db, creds.Name, creds.Email, string(hashedPassword))
		if err != nil {
			fmt.Println(err)
			creds.ErrorE = true
			tmpl.Execute(w, creds)
			return
		}

		http.Redirect(w, r, "/signin", 301)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

}
