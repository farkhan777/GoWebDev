package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":7777", mux)
}

func user(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(res, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(res, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	HandleError(res, err)
}

func about(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	HandleError(res, err)
}

func contact(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "contact.gohtml", nil)
	HandleError(res, err)
}

func apply(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "apply.gohtml", nil)
	HandleError(res, err)
}

func applyProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "applyProcess.gohtml", nil)
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
