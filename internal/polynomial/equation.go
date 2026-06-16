package polynomial

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/tmazitov/computorv1/internal/parsing"
)

type Equation struct {
	scalar    parsing.ScalarMap
	maxDegree int
}

func NewEquation(scalarMap parsing.ScalarMap) (*Equation, error) {

	maxDegree := scalarMap.GetMaxDegree()

	return &Equation{
		scalar: scalarMap,
		maxDegree: maxDegree,
	}, nil
}

func (e Equation) Degree() int {
	return e.maxDegree
}

func (e Equation) Scalar() parsing.ScalarMap {
	return e.scalar
}

func (e Equation) ToString() string {
	var degrees []int

	for degree := range e.scalar {
		degrees = append(degrees, degree)
	}
	sort.Ints(degrees)

	var result string
	for i, degree := range degrees {
		scalar := e.scalar[degree]
		coeff := strconv.FormatFloat(float64(scalar), 'f', -1, 32)

		if i == 0 {
			result += fmt.Sprintf("%s * X^%d", coeff, degree)
		} else if scalar < 0 {
			result += fmt.Sprintf(" - %s * X^%d", coeff[1:], degree)
		} else {
			result += fmt.Sprintf(" + %s * X^%d", coeff, degree)
		}
	}

	return result + " = 0"
}
