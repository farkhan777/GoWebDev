package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
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
	http.SetCookie(w, c)

	var u user
	if un, ok := dbSession[c.Value]; ok {
		fmt.Println(dbSession[c.Value])
		fmt.Println("U atas: ", u)
		u = dbUsers[un]
		fmt.Println("U bawah: ", u)
	}
	fmt.Println(u)
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSession[c.Value]
	_, ok := dbUsers[un]

	return ok
}
