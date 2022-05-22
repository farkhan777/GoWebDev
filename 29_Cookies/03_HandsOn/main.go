package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	//http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	//count, err := strconv.ParseInt(cookie.Value, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	//cookie.Value = strconv.FormatInt(count, 10)
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, cookie.Value)
}

//func read(w http.ResponseWriter, r *http.Request)  {
//	cookie, err := r.Cookie("my-cookie")
//
//	if err == http.ErrNoCookie {
//		cookie = &http.Cookie{
//			Name:  "my-cookie",
//			Value: "0",
//			Path:  "/read",
//		}
//	}
//
//	count, err := strconv.Atoi(cookie.Value)
//	//count, err := strconv.ParseInt(cookie.Value, 10, 64)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	count++
//	cookie.Value = strconv.Itoa(count)
//	//cookie.Value = strconv.FormatInt(count, 10)
//	http.SetCookie(w, cookie)
//
//	fmt.Fprintln(w, cookie.Value)
//}
