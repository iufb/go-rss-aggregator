package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	ah := headers.Get("Authorization")
	if ah == "" {
		return "", errors.New("no authorization info found.")
	}
	vals := strings.Split(ah, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization info.")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of  authorization info.")
	}
	return vals[1], nil
}
