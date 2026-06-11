package polynomial

import (
	"regexp"
	"strconv"
	"strings"
)

func getMaxDegree(scalarMap map[int]float32) int {
	var maxDegree int = -1

	for degree := range scalarMap {
		if degree > maxDegree {
			maxDegree = degree
		}
	}

	return maxDegree
}

func extractFromString(raw string) (map[int]float32, error) {

	sides := strings.Split(raw, " = ")

	var (
		scalarMap    map[int]float32 = map[int]float32{}
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

	return scalarMap, nil
}
