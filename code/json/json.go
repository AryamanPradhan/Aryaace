package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_Name"`
	LastName  string `json:"last_Name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
	[
		{
			"first_Name": "Clark",
			"last_Name": "Kent",
			"hair_color":"Black",
			"has_dog": true
		},
		{
			"first_Name": "John",
			"last_Name": "Wick",
			"hair_color":"Brown",
			"has_dog": false

		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Print("Error unmarshalling json", err)
	}
	log.Printf("Unmarshalled %v", unmarshalled)

	//Write	json from struct

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Aryaman"
	m2.LastName = "Pradhan"
	m2.HairColor = "Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Print("Error marshalling", err)
	}

	fmt.Println(string(newJson))

}
