package greetings

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"tutorial/greetings/errs"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("hello: %w", errs.NewEmptyNameError())
	}

	format := randomFormat()
	message := fmt.Sprintf(format, name)

	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
	}

	size := len(formats)
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))

	return formats[n.Int64()]
}
