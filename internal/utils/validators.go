package utils

import (
	"net/http"

	"golang.org/x/exp/slices"
)

var HTTP_METHODS = []string{
	http.MethodGet,
	http.MethodPost,
}

func ValidateHttpMethodName(method string) bool {
	return slices.Contains(HTTP_METHODS, method)
}
