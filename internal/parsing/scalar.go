package parsing

import (
	"regexp"
	"errors"
	"strconv"
	"strings"
)

type ScalarMap map[int]float32

func (m ScalarMap) GetMaxDegree() int {
	var maxDegree int = -1

	for degree, scalar := range m {
		if scalar != 0 && degree > maxDegree {
			maxDegree = degree
		}
	}

	return maxDegree
}

func NewScalarMap(raw string) (ScalarMap, error) {

	if len(raw) == 0 {
		return nil, errors.New("computorv1 parsing error: empty equation")
	}

	sides := strings.Split(raw, " = ")

	var (
		scalarMap    ScalarMap = map[int]float32{}
		tempSign     bool
		tempScalar   float32
		scalarRegexp *regexp.Regexp = regexp.MustCompile(`^[+-]?(\d+(\.\d+)?|\.\d+)$`)
		degreeRegexp *regexp.Regexp = regexp.MustCompile(`^X\^\d+$`)
	)

	for sideIndex, side := range sides {

		parts := strings.Split(side, " ")

		for _, part := range parts {
			if part == "-" {
				tempSign = true
				continue
			}

			if scalarRegexp.MatchString(part) {

				scalar, err := strconv.ParseFloat(part, 64)
				if err != nil {
					return nil, err
				}

				tempScalar = float32(scalar)
			}

			if degreeRegexp.MatchString(part) {
				degree, err := strconv.Atoi(part[2:])
				if err != nil {
					return nil, err
				}

				if tempSign {
					tempScalar *= -1
				}
				if sideIndex != 0 {
					tempScalar *= -1
				}

				if _, ok := scalarMap[degree]; !ok {
					scalarMap[degree] = 0
				}

				scalarMap[degree] += tempScalar
				tempScalar = 0
				tempSign = false
			}
		}
	}

	if len(scalarMap) == 0 {
		return nil, errors.New("computorv1 parsing error: empty equation")
	}

	return scalarMap, nil
}


func (m ScalarMap) IsEmpty() bool {
	for _, scalar := range m {
		if scalar != 0 {
			return false
		}
	}
	return true
}