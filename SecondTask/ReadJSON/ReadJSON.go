package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Friend struct {
	Fname  string `json:"fname"`
	Sname  string `json:"sname"`
	Gender string `json:"gender"`
	Height string `json:"hight"`
}

func main() {
	content, err := ioutil.ReadFile("/Users/u746582/IdeaProjects/GoTaskNr2/SecondTask/ReadJSON/Friends.json")
	if err != nil {
		log.Fatal(err)
	}

	var friends []Friend
	err2 := json.Unmarshal(content, &friends)
	if err2 != nil {
		log.Fatal(err)
	}

	for _, x := range friends {
		fmt.Println(x.Fname, x.Sname)
	}
}
