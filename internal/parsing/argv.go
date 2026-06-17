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

	if args[1] == "" {
		return "", errors.New("computorv1 parsing error: empty equation")
	}

	return args[1], nil
}
