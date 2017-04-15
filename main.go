package main

import (
	"fmt"

	"github.com/rojoherrero/hello/models"

	"github.com/rojoherrero/hello/functions"
)

var (
	v1 = 428
	v2 = 27
)

var otherArray [10]int

func main() {

	people := []models.Person{models.Person{Name: "Sergio"}, models.Person{Name: "Pili"}, models.Person{Name: "Alberto"}}

	for index, item := range people {
		fmt.Printf("On position %d is %s\n", index, item.Name)
	}

	result, modulo := functions.Divide(v1, v2)

	fmt.Printf("%[1]d divided by %[2]d has a result of %[3]d with reminder of %[4]d\n", v1, v2, result, modulo)
	fmt.Printf("Probe -> (result*v2)+modulo = %d\n", (result*v2)+modulo)

}
