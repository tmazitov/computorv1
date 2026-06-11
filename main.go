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
}
