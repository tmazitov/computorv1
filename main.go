package main

import (
	"fmt"
	"log"

	"github.com/tmazitov/computorv1/internal/parsing"
	"github.com/tmazitov/computorv1/internal/polynomial"
)

func main() {

	raw, err := parsing.GetEquation()
	if err != nil {
		log.Fatal(err)
	}

	equation, err := polynomial.NewEquation(raw)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reduced form: %s\n", equation.ToString())

	equationDegree := equation.MaxDegree()
	if equationDegree > 0 {
		fmt.Printf("Polynomial degree: %d\n", equationDegree)
	} else if !equation.IsEmpty() {
		fmt.Println("No solution.")
		return
	} else {
		fmt.Println("Any real number is a solution.")
		return
	}

	
}
