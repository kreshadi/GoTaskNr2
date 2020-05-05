package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//fmt.Println("Hi")
	tpl, err := template.ParseFiles("/Users/u746582/IdeaProjects/GoTaskNr2/SecondTask/tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
