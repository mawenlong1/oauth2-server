package oauth

import (
	"errors"
	"oauth2-server/models"
	"sort"
	"strings"
)

var (
	// ErrInvalidScope ..
	ErrInvalidScope = errors.New("oauth:Invalid scope")
)

// GetScope ..
func (s *Service) GetScope(requestScope string) (string, error) {
	if requestScope == "" {
		return s.GetDefaultScope(), nil
	}
	if s.ScopeExists(requestScope) {
		return requestScope, nil
	}
	return "", ErrInvalidScope
}

// GetDefaultScope ..
func (s *Service) GetDefaultScope() string {
	var scopes []string
	s.db.Model(new(models.OauthScope)).Where("is_default=?", true).Pluck("scope", &scopes)
	sort.Strings(scopes)
	return strings.Join(scopes, " ")
}

// ScopeExists ..
func (s *Service) ScopeExists(requestScope string) bool {
	scopes := strings.Split(requestScope, " ")
	var count int
	s.db.Model(new(models.OauthScope)).Where("scope in (?)", scopes).Count(&count)
	return count == len(scopes)
}
