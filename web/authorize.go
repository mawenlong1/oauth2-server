package web

import (
	"errors"
	"net/http"
)

var (
	// ErrIncorrectResponseType ..
	ErrIncorrectResponseType = errors.New("web:Response type not one of token or code")
)

func (s *Service) authorizeForm(w http.ResponseWriter, r *http.Request) {

}
