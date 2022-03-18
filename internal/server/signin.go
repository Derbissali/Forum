package server

import (
	"database/sql"
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (handle *Dbs) Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auther")
	tmpl, err := template.ParseFiles(cfg.HTML + "login.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	errU := internal.User{}
	_, err = r.Cookie("session")
	if err == nil {
		log.Println("signin")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		tmpl.Execute(w, errU)
		return
	case "POST":
		creds := &internal.User{
			Name:     r.FormValue("login"),
			Password: r.FormValue("password"),
		}

		result := handle.Db.QueryRow(`SELECT "password" from "user" WHERE name=$1`, creds.Name)

		ourPerson := internal.User{}
		err := result.Scan(&ourPerson.Password)
		if err != nil {
			// If an entry with the username does not exist, send an "Unauthorized"(401) status
			if err == sql.ErrNoRows {
				fmt.Println("wrong login")
				errU.ErrorL = true
				tmpl.Execute(w, errU)
				return
			}
			// If the error is of any other type, send a 500 status
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("error")
			return
		}
		//	fmt.Println(ourPerson.Password, creds.Password)
		if err = bcrypt.CompareHashAndPassword([]byte(ourPerson.Password), []byte(creds.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			fmt.Println("wrong password")
			errU.ErrorL = true
			tmpl.Execute(w, errU)
			return
		}
		//	fmt.Println(creds.Name)
		c, err := r.Cookie("session")

		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		a := database.SelectUserID(handle.Db, creds.Name)
		err = database.AddSession(handle.Db, c.Value, a)
		if err != nil {
			b := database.SelectUser2(handle.Db, creds.Name)

			errU.ID = b
			fmt.Println(errU, "-------------s")
			database.DeleteSession(handle.Db, errU.ID)
			errU.Session = true
			http.SetCookie(w, c)
			database.AddSession(handle.Db, c.Value, a)
			fmt.Println("authed user", errU)
			http.Redirect(w, r, "/", 301)
			return

		}

		http.Redirect(w, r, "/", 301)
		return
	default:
		fmt.Println("dleert")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl.Execute(w, errU)

}
