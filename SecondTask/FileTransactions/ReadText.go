package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("/Users/u746582/IdeaProjects/GoTaskNr2/SecondTask/text1.txt", "/Users/u746582/IdeaProjects/GoTaskNr2/SecondTask/text2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n----------")
	tpl.ExecuteTemplate(os.Stdout, "text1.txt", nil)
	fmt.Println("\n----------")
	tpl.ExecuteTemplate(os.Stdout, "text2.txt", "Kamellia")
	fmt.Println("\n----------")

}
