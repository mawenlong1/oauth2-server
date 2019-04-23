package session_test

import (
	"github.com/stretchr/testify/assert"
	"oauth2-server/session"
)

func (suite *SessionTestSuite) TestService() {
	var (
		userSession *session.UserSession
		err         error
	)
	// 没有启动session
	userSession, err = suite.service.GetUserSession()
	assert.Nil(suite.T(), userSession)
	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), session.ErrSessionNotStarted, err)
	}
	// 启动session
	err = suite.service.StartSession()
	assert.Nil(suite.T(), err)
	// 清除session
	err = suite.service.ClearUserSession()
	assert.Nil(suite.T(), err)
	// 获取userSerssion,由于session没有完成返回error
	userSession, err = suite.service.GetUserSession()
	assert.Nil(suite.T(), userSession)
	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), "User session type assertion error", err.Error())
	}
	// 设置用户session
	err = suite.service.SetUserSession(&session.UserSession{
		ClientID:     "test_cilent",
		Username:     "test@username",
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	})
	assert.Nil(suite.T(), err)
	// 	获取session
	userSession, err = suite.service.GetUserSession()
	assert.Nil(suite.T(), err)
	if assert.NotNil(suite.T(), userSession) {
		assert.Equal(suite.T(), "test_cilent", userSession.ClientID)
		assert.Equal(suite.T(), "test@username", userSession.Username)
		assert.Equal(suite.T(), "test_access_token", userSession.AccessToken)
		assert.Equal(suite.T(), "test_refresh_token", userSession.RefreshToken)
	}
	// 清除session
	err = suite.service.ClearUserSession()
	assert.Nil(suite.T(), err)
	// 获取session
	userSession, err = suite.service.GetUserSession()
	assert.Nil(suite.T(), userSession)
	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), "User session type assertion error", err.Error())
	}
}
