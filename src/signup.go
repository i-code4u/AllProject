package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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

	type em struct {
		Email string
	}
	check_email := em{r.Form.Get("email")}
	var result em
	err := collection.FindOne(context.TODO(), check_email).Decode(&result)
	if result.Email == r.Form.Get("email") {
		message := "Already registered"
		err = tpl.ExecuteTemplate(w, "login.html", message)
		if err != nil {
			fmt.Println("error in index", err)
		}
	} else {
		insertResult, err := collection.InsertOne(context.TODO(), dt1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		err = tpl.ExecuteTemplate(w, "select_room.html", dt1)
		if err != nil {
			fmt.Println("error in index", err)
		}
		err = client.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}
}
