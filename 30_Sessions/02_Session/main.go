package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	Username string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user, ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	// process from submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
