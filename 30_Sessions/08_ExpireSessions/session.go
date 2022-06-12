package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSession[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSession[c.Value] = s
		u = dbUser[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSession[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSession[c.Value] = s
	}
	_, ok = dbUser[s.un]

	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSession {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSession, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()
}

func showSessions() {
	fmt.Println("********")
	for k, v := range dbSession {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
