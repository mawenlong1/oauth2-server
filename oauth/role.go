package oauth

import (
	"errors"
	"oauth2-server/models"
)

var (
	//ErrRoleNotFound ...
	ErrRoleNotFound = errors.New("Role not found")
)

//FindRoleByID ...
func (s *Service) FindRoleByID(id string) (*models.OauthRole, error) {
	role := new(models.OauthRole)
	if s.db.Where("id=?", id).First(role).RecordNotFound() {
		return nil, ErrRoleNotFound
	}
	return role, nil
}
