package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	type Person struct {
		Name  string
		Phone string
	}
	session, err := mgo.Dial("localhost")
	//fmt.Println(session)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("demo").C("people")
	/*err = c.Insert(&Person{"Shivmani", "+1234"},
		&Person{"Shambhoo", "+3427"})
	//	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}*/

	result := Person{}
	err = c.Find(bson.M{"name": "Shivmani"}).One(&result)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println("Phone:", result.Phone)
}
