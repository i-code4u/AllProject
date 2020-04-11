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
	mux.GET("/login", login)
	mux.POST("/login", login)
	mux.POST("/passdata", passdata)
	mux.ServeFiles("/public/*filepath", http.Dir("../public/"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println("error in login")
	}
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("error in index")
	}
}
func passdata(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	type newreg struct {
		Name     string
		Email    string
		Password string
		Contact  string
	}
	dt1 := newreg{r.Form.Get("name"), r.Form.Get("email"), r.Form.Get("password"), r.Form.Get("contact")}
	err := tpl.ExecuteTemplate(w, "passdata.html", dt1)
	if err != nil {
		fmt.Println("error in index", err)
	}
}
