package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println("error in login")
	}
}

func signIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	type newreg struct {
		Email    string
		Password string
		//name     string
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

	} else {
		message := "Invalid email or Password!!!"
		err = tpl.ExecuteTemplate(w, "login.html", message)
		if err != nil {
			fmt.Println("error in index", err)
		}
	}
}
