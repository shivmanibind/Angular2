package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	//ID   string `json:"id"`
	Name string `json:"name"`
	//Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var people []Person

func getDataEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	email := vars["email"]
	fmt.Print("Hee", email)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydb").C("userdetails")

	result := Person{}
	err = c.Find(bson.M{"email": email}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(people)
	people = append(people, Person{Name: result.Name, Email: result.Email, Password: result.Password})
}

func postDataEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	//	person.ID = params["id"]
	//json.NewEncoder(w).Encode(people)
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydb").C("userdetails")
	err = c.Insert(&Person{person.Name, person.Email, person.Password})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("In post")
}

func updateDataEndPoint(w http.ResponseWriter, req *http.Request) {

}

func deleteDataEndPoint(w http.ResponseWriter, req *http.Request) {

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/data/{email}", getDataEndPoint).Methods("GET")
	router.HandleFunc("/data/user/{id}", postDataEndPoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))

}
