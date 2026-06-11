package polynomial

import (
	"fmt"
	"sort"
	"strconv"
)

type Equation struct {
	scalar    map[int]float32
	maxDegree int
}

func NewEquation(raw string) (*Equation, error) {

	scalar, err := extractFromString(raw)
	if err != nil {
		return nil, err
	}

	maxDegree := getMaxDegree(scalar)

	return &Equation{
		scalar:    scalar,
		maxDegree: maxDegree,
	}, nil
}

func (e Equation) MaxDegree() int {
	return e.maxDegree
}

func (e Equation) IsEmpty() bool {
	for _, scalar := range e.scalar {
		if scalar != 0 {
			return false
		}
	}
	return true
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
