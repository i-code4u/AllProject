package main

import (
	"context"
	"fmt"
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
func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println("error in login")
	}
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("error in index")
	}

}
func signUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	type newreg struct {
		Name     string
		Email    string
		Password string
		Contact  string
	}
	collection := client.Database("data").Collection("registration") // collection
	dt1 := newreg{r.Form.Get("name"), r.Form.Get("email"), r.Form.Get("password"), r.Form.Get("contact")}
	///////////////
	insertResult, err := collection.InsertOne(context.TODO(), dt1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	///////////////
	err = tpl.ExecuteTemplate(w, "passdata.html", dt1)
	if err != nil {
		fmt.Println("error in index", err)
	}
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
func signIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	type newreg struct {
		Email    string
		Password string
	}
	collection := client.Database("data").Collection("registration") // collection
	dt1 := newreg{r.Form.Get("email"), r.Form.Get("password")}
	///////////////
	var result newreg
	err := collection.FindOne(context.TODO(), dt1).Decode(&result)
	fmt.Printf("Found a single document: %+v\n", result)
	if result.Email == r.Form.Get("email") && result.Password == r.Form.Get("password") {
		err = tpl.ExecuteTemplate(w, "passdata.html", nil)
		if err != nil {
			fmt.Println("error in index", err)
		}
		err = client.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")

	} else {
		message := "Invalid email or Password!!!"
		err = tpl.ExecuteTemplate(w, "login.html", message)
		if err != nil {
			fmt.Println("error in index", err)
		}
	}
}
