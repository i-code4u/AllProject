package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func room1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type new struct {
		Count          string
		Organiser      string
		Date           string
		Start_time     string
		End_time       string
		Attendee_email []string
		User_form      string
		Status         string
	}
	var results []*new
	findOptions := options.Find()
	collection := client.Database("data").Collection("addedRoom1Meeting")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for cur.Next(context.TODO()) {
		count++
		// create a value into which the single document can be decoded
		var elem new
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count)
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	i := 0
	for i < count {
		dbtime1 := results[i].Date
		dbday, _ := strconv.Atoi(dbtime1[3:5])
		dbmin1 := results[i].Start_time
		dbh1, _ := strconv.Atoi(dbmin1[0:2])
		dbtime2 := results[i].End_time
		dbh2, _ := strconv.Atoi(dbtime2[0:2])
		dbmin2 := results[i].Start_time
		dbm1, _ := strconv.Atoi(dbmin2[3:5])
		dbdate := results[i].End_time
		dbm2, _ := strconv.Atoi(dbdate[3:5])
		fmt.Println("dbtime", i, dbh1, dbm1, dbday)
		time := time.Now()
		hour := time.Hour()
		minute := time.Minute()
		day := time.Day()
		fmt.Println("currunt ", i, hour, minute, day)
		if dbh1 == hour && dbday == day && dbm1 == minute {
			results[i].Status = "Ongoing"
		} else if dbday == day && hour == dbh1 && minute > dbm1 && dbm2 > minute {
			results[i].Status = "Ongoing"

		} else if dbday == day && minute <= dbm1 && minute <= dbm2 && dbh2 <= hour {
			results[i].Status = "Ongoing"
		}
		i++
	}
	err = tpl.ExecuteTemplate(w, "room1.html", results)
	if err != nil {
		fmt.Println("error in login", err)
	}
}
