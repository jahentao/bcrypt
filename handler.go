package function

import (
	"net/http"

	"github.com/openfaas-incubator/go-function-sdk"
	"golang.org/x/crypto/bcrypt"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	hash, err := bcrypt.GenerateFromPassword(req.Body, bcrypt.DefaultCost)
	if err != nil {
		return handler.Response{
			Body:       []byte(""),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return handler.Response{
		Body:       hash,
		StatusCode: http.StatusOK,
	}, err
}
