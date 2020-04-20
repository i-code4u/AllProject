package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func addmeeting(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "AddMeeting.html", nil)
	if err != nil {
		fmt.Println("error in addmeeting", err)
	}
}
