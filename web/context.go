package web

import (
	"errors"
	"github.com/gorilla/context"
	"net/http"
	"oauth2-server/models"
	"oauth2-server/session"
)

type contextKey int

const (
	sessionServiceKey contextKey = 0
	clientKey         contextKey = 1
)

var (
	// ErrSessionServiceNotPresent ..
	ErrSessionServiceNotPresent = errors.New("web:Session service not present in the request context")
	// ErrClientNotPresent ...
	ErrClientNotPresent = errors.New("web:Client not present in the request context")
)

func getSessionService(r *http.Request) (session.ServiceInterface, error) {
	val, ok := context.GetOk(r, sessionServiceKey)
	if !ok {
		return nil, ErrSessionServiceNotPresent
	}
	sessionService, ok := val.(session.ServiceInterface)
	return sessionService, nil
}
func getClient(r *http.Request) (*models.OauthClient, error) {
	val, ok := context.GetOk(r, clientKey)
	if !ok {
		return nil, ErrClientNotPresent
	}
	client, ok := val.(*models.OauthClient)
	if !ok {
		return nil, ErrClientNotPresent
	}
	return client, nil
}
