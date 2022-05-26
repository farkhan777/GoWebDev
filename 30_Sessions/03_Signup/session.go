package main

import "net/http"

func getUser(r *http.Request) user {
	var u user

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exist already, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
