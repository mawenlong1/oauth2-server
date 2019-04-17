package oauth

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	"oauth2-server/models"
	"oauth2-server/util"
	"oauth2-server/util/password"
	"strings"
	"time"
)

var (
	//ErrClientNotFound ..
	ErrClientNotFound = errors.New("Client not found")
	//ErrInvalidClientSecret ..
	ErrInvalidClientSecret = errors.New("Invalid client Secret")
	//ErrClientIDTaken ..
	ErrClientIDTaken = errors.New("ClientID exists")
)

//ClientExists ...
func (s *Service) ClientExists(clientID string) bool {
	_, err := s.FindRoleByID(clientID)
	return err == nil
}

//FindClientByClientID ..
func (s *Service) FindClientByClientID(clientID string) (*models.OauthClient, error) {
	client := new(models.OauthClient)
	notFound := s.db.Where("client_key=LOWER(?)", clientID).First(client).RecordNotFound()
	if notFound {
		return nil, ErrClientNotFound
	}
	return client, nil
}

//CreateClient ..
func (s *Service) CreateClient(clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(s.db, clientID, secret, redirectURI)
}

//CreateClientTx ...
func (s *Service) CreateClientTx(tx *gorm.DB, clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(tx, clientID, secret, redirectURI)
}

func (s *Service) createClientCommon(db *gorm.DB, clientID, secret, redirectURI string) (*models.OauthClient, error) {
	if s.ClientExists(clientID) {
		return nil, ErrClientIDTaken
	}
	secretHash, err := password.HashPassword(secret)
	if err != nil {
		return nil, err
	}
	client := &models.OauthClient{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientKey:    strings.ToLower(clientID),
		ClientSecret: string(secretHash),
		RedirectURI:  util.StringOrNull(redirectURI),
	}
	if err := db.Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

//AuthClient ...
func (s *Service) AuthClient(clientID, secret string) (*models.OauthClient, error) {
	client, err := s.FindClientByClientID(clientID)
	if err != nil {
		return nil, ErrClientNotFound

	}
	if password.VerifyPassword(client.ClientSecret, secret) != nil {
		return nil, ErrInvalidClientSecret
	}
	return client, nil
}
