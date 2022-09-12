package adapter

import (
	"errors"
	"strings"
)

type StringService struct{}

func (sc StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (sc StringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("empty string")
