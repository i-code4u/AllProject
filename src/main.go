package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var client, _ = mongo.Connect(context.TODO(), clientOptions)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../views/*"))
}
func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/login", login)
	mux.POST("/login", login)
	mux.POST("/signUp", signUp)
	mux.POST("/signIn", signIn)
	mux.ServeFiles("/public/*filepath", http.Dir("../public/"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
