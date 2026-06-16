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

	scalarMap, err := parsing.NewScalarMap(raw)
	if err != nil {
		log.Fatal(err)
	}

	equation, err := polynomial.NewEquation(scalarMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reduced form: %s\n", equation.ToString())

	equationDegree := equation.Degree()
	if equation.Scalar().IsEmpty() {
		fmt.Println("Any real number is a solution.")
		return
	} else if equationDegree == 0{
		fmt.Println("No solution.")
		return
	}

	fmt.Printf("Polynomial degree: %d\n", equationDegree)
	if equationDegree > 2 {
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		return
	}	
	
	roots := equation.Solve()
	if len(roots) == 0 {
		fmt.Println("Discriminant is strictly negative, no real solution.")
	} else if len(roots) == 1 {
		fmt.Printf("Discriminant is zero, the solution is:\n%.6g\n", roots[0])
	} else {
		fmt.Printf("Discriminant is strictly positive, the two solutions are:\n%.6g\n%.6g\n", roots[0], roots[1])
	}
}
