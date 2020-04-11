package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../views/*"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.POST("/login", login)
	http.Handle("./public/", http.StripPrefix("./public", http.FileServer(http.Dir("./public")))) //??????
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	err := tpl.ExecuteTemplate(w, "login.html", 3)
	if err != nil {
		fmt.Println("error in login")
	}
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	http.Handle("./public/", http.StripPrefix("./public/", http.FileServer(http.Dir("./public/stylesheet"))))
	r.ParseForm()
	err := tpl.ExecuteTemplate(w, "index.html", 3)
	if err != nil {
		fmt.Println("error in index")
	}
}
