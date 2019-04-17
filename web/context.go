package web

import (
	"errors"
	"github.com/gorilla/context"
	"net/http"
	"oauth2-server/session"
)

type contextKey int

const (
	sessionServiceKey contextKey = 0
)

var (
	ErrSessionServiceNotParent = errors.New("")
)

func getSessionService(r *http.Request) (session.ServiceInterface, error) {
	val, ok := context.GetOk(r, sessionServiceKey)
	if !ok {
		return nil, ErrSessionServiceNotParent
	}
	sessionService, ok := val.(session.ServiceInterface)
	return sessionService, nil
}
