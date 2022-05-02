package main

import (
	"fmt"
	"github.com/sajitron/everyday-go/get-json/astros"
	"log"
)

func main() {
	url := "http://api.open-notify.org/astros.json"

	people, err := astros.GetAstros(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%d people found in space.\n", people.Number)

	for _, p := range people.Person {
		fmt.Printf("Let's wave to: %s\n", p.Name)
	}
}
