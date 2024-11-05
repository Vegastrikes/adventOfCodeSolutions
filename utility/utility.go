package utility

import (
	"errors"
	"os"
)

func GetInput(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := string(dat)

	if len(result) == 0 {
		return "", errors.New("empty file")
	}

	return result, nil
}
