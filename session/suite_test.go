package session_test

import (
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"oauth2-server/config"
	"oauth2-server/session"
	"testing"
)

// SessionTestSuite ...
type SessionTestSuite struct {
	suite.Suite
	cfg     *config.Config
	service *session.Service
}

func (suite *SessionTestSuite) SetupSuite() {
	suite.cfg = config.NewDefaultConfig()

	session.StorageSessionName = "test_session"
	session.UserSessionKey = "test_user"
	r, err := http.NewRequest("GET", "http://1.2.3.4/foo/bar", nil)
	assert.NoError(suite.T(), err, "request setup should not get an error")
	w := httptest.NewRecorder()
	suite.service = session.NewService(suite.cfg, sessions.NewCookieStore([]byte(suite.cfg.Session.Secret)))
	suite.service.SetSessionService(r, w)
}

func TestSessionTestSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}
