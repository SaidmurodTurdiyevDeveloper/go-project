package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//we can change name when we are changing json
	Name     string `json:"coursename"`
	Price    int
	Platform string
	//we can remove this value when we are changing json with -
	Password string `json:"-"`
	//we can remove this value will be nil when we are changing json
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json Parsing Lesson")
	//EncodingJson()
	DecodingJson()
}

func EncodingJson() {
	courses := []course{
		{"ReactJs BootCamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern BootCamp", 199, "LearnCodeOnline.in", "bsd654", []string{"js"}},
		{"Angular BootCamp", 322, "LearnCodeOnline.in", "fdgr23", nil},
	}
	/*
		//change json simple
		myJson, err := json.Marshal(courses)
	*/
	//change json with struct
	myJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	defer func() {
		result := recover()
		if result != nil {
			fmt.Println(result)
		}
	}()
	fmt.Printf("%s\n", myJson)
}

func DecodingJson() {
	courseJson := []byte(`
	{
		"coursename": "React JS BootCamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
		  "web-dev",
		  "js"
		]
	}
	`)

	chracterJson := []byte(`
	{
		"id": 1,
		"name": "Rick Sanchez",
		"status": "Alive",
		"species": "Human",
		"type": "",
		"gender": "Male",
		"origin": {
		  "name": "Earth (C-137)",
		  "url": "https://rickandmortyapi.com/api/location/1"
		}
	}
	`)
	fmt.Println("\nCourse data")
	var lcoCourse course

	checkValid := json.Valid(courseJson)
	if checkValid {
		fmt.Println("JsonCourse is valid")
		json.Unmarshal(courseJson, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JsonCourse is not valid")
	}
	fmt.Println("\nCharacter data")

	var myCharcater character

	checkValidCharacter := json.Valid(chracterJson)
	if checkValidCharacter {
		fmt.Println("JsonCharacter is valid")
		json.Unmarshal(chracterJson, &myCharcater)
		fmt.Printf("%#v\n", myCharcater)
	} else {
		fmt.Println("JsonCharacter is not valid")
	}

	fmt.Println("\nCourse Map")
	var myCourseMap map[string]interface{}
	json.Unmarshal(courseJson, &myCourseMap)

	for k, v := range myCourseMap {
		fmt.Printf("Key \"%v\" Value \"%v\" Type \"%T\"\n", k, v, v)
	}

	fmt.Println("\nCharacter Map")
	var myCharactereMap map[string]interface{}
	json.Unmarshal(chracterJson, &myCharactereMap)
	for k, v := range myCharactereMap {
		if vM, ok := v.(map[string]interface{}); ok {
			fmt.Println("")
			fmt.Printf("Key \"%v\"\nThis %v values:\n", k, k)
			for kM, vMap := range vM {
				fmt.Printf("Key \"%v\" Value \"%v\" Type \"%T\"\n", kM, vMap, vMap)
			}
			fmt.Println("")
		} else {
			fmt.Printf("Key \"%v\" Value \"%v\" Type \"%T\"\n", k, v, v)
		}
	}

}

type character struct {
	Id      int
	Name    string
	Status  string
	Species string
	Type    string
	Gender  string
	Origin  orign
}
type orign struct {
	Name string
	Url  string
}
