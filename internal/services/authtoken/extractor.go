package authtoken

import (
	"errors"
	"strings"
)

const bearerTemplate = "Bearer "

var ErrInvalidToken = errors.New("token is not valid")

func ExtractToken(authHeader string) (string, error) {
	if len(authHeader) == 0 {
		return "", errors.New("auth header is empty")
	}

	splitHeader := strings.Split(authHeader, bearerTemplate)
	if len(splitHeader) != 2 {
		return "", ErrInvalidToken
	}

	token := strings.TrimSpace(splitHeader[1])
	if len(token) == 0 {
		return "", ErrInvalidToken
	}

	return token, nil
}
