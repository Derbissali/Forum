package server

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"net/http"
	"text/template"
)

func (handle *Dbs) Posting(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(cfg.HTML + "addpost.html")
	r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Println("CASE - POST")
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	M := internal.Forum{}
	rows, e := handle.Db.Query(`SELECT "name" FROM "category" ORDER BY "name"`)
	if e != nil {
		return
	}
	for rows.Next() {
		cat1 := internal.Category{}
		e = rows.Scan(&cat1.Name)
		if e != nil {
			return
		}
		M.Category = append(M.Category, cat1)

	}
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		fmt.Println("qwe1")
		http.Redirect(w, r, "/signin", 301)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)

	n := M.User.ID

	var Cat []string
	r.ParseForm()
	Cat = r.Form["category"]
	fmt.Println("cat", Cat)
	if len(r.FormValue("postN")) > 50 || len(r.FormValue("postB")) > 2000 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if len(r.FormValue("postN")) == 0 || len(r.FormValue("postB")) == 0 {
		M.Post.TitBodNull = true
		tmpl.Execute(w, M)
		return
	}

	fmt.Println(r.FormValue("postN"))
	file, handler, err := r.FormFile("myFile")
	var fileName string

	if err == nil {
		dst, err := os.Create(fmt.Sprintf("assets/temp-images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()
		fileName = strings.TrimPrefix(dst.Name(), "assets/temp-images/")
		if len(fileName) == 0 {
			fileName = "1"
		}
		fmt.Println(fileName) // Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//func() {}
	}

	if len(Cat) > 0 {
		b := 0
		stmt := handle.Db.QueryRow(`INSERT INTO post (name, body, user_id, image, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?) RETURNING id`, r.FormValue("postN"), r.FormValue("postB"), n, fileName, 0, 0)
		stmt.Scan(&b)
		for _, i := range Cat {
			a := 0

			stmt1 := handle.Db.QueryRow(`SELECT "id" FROM "category" WHERE category.name=?`, i)
			stmt1.Scan(&a)
			if a == 0 {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			_, err = handle.Db.Exec(`INSERT INTO category_post (category_id, post_id) VALUES (?, ?) `, a, b)
			if err != nil {
				fmt.Println(err)
			}

		}

	} else {
		M.Post.CategoryNull = true
		tmpl.Execute(w, M)
		return
	}
	http.Redirect(w, r, "/", 301)
	return
}
func (handle *Dbs) AddPost(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(cfg.HTML + "addpost.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// return
	// if r.URL.Path != "/createpost" {
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	// 	return
	// }
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	M := internal.Forum{}

	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err, "errro")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	M.User, err = database.ReadByUuid(handle.Db, c.Value)

	M.Category = database.SelectCat(handle.Db)

	tmpl.Execute(w, M)
	return

}
