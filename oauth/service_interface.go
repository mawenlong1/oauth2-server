package oauth

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"oauth2-server/config"
	"oauth2-server/models"
	"oauth2-server/session"
	"oauth2-server/util/routes"
)

//ServiceInterface ...
type ServiceInterface interface {
	GetConfig() *config.Config
	RestrictToRoles(allowedRoles ...string)
	IsRoleAllowed(role string) bool

	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)

	FindRoleByID(id string) (*models.OauthRole, error)

	ClientExists(clientID string) bool
	FindClientByClientID(clientID string) (*models.OauthClient, error)
	CreateClient(clientID, secret, redirectURI string) (*models.OauthClient, error)
	CreateClientTx(tx *gorm.DB, clientID, secret, redirectURI string) (*models.OauthClient, error)
	AuthClient(clientID, secret string) (*models.OauthClient, error)

	UserExists(username string) bool
	FindUserByUserName(username string) (*models.OauthUser, error)
	CreateUser(roleID, username, password string) (*models.OauthUser, error)
	CreateUserTx(tx *gorm.DB, roleID, username, password string) (*models.OauthUser, error)
	SetPassword(user *models.OauthUser, password string) error
	SetPasswordTx(tx *gorm.DB, user *models.OauthUser, password string) error
	UpdateUsername(user *models.OauthUser, username string) error
	UpdateUsernameTx(tx *gorm.DB, user *models.OauthUser, username string) error
	AuthUser(username, thePassword string) (*models.OauthUser, error)

	GetScope(requestScope string) (string, error)
	GetDefaultScope() string
	ScopeExists(requestScope string) bool

	Login(client *models.OauthClient, user *models.OauthUser, scope string) (*models.OauthAccessToken, *models.OauthRefereshToken, error)
	GrantAuthorizationCode(client *models.OauthClient, user *models.OauthUser, expiresIn int, redirectURI, scope string) (*models.OauthAuthorizationCode, error)
	GrandAccessToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthRefereshToken, error)
	GetOrCreateRefreshToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthRefereshToken, error)
	GetValidRefreshToken(token string, client *models.OauthClient) (*models.OauthRefereshToken, error)
	Authenticate(token string) (*models.OauthAccessToken, error)

	ClearUserTokens(userSession *session.UserSession)
	Close()
}
