package parsing

import (
	"errors"
	"os"
)

func GetEquation() (string, error) {
	args := os.Args
	if len(args) == 1 {
		return "", nil
	}

	if len(args) != 2 {
		return "", errors.New("computorv1 parsing error: invalid number of arguments")
	}
	return args[1], nil
}
