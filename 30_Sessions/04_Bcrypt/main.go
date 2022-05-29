package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	Username string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			for k, e := range dbUsers {
				fmt.Println("Key: ", k, "\nElement: ", e)
			}

			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// store user dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u = user{un, bs, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}
