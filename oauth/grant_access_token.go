package oauth

import (
	"oauth2-server/models"
	"time"
)

// GrandAccessToken ...
func (s *Service) GrandAccessToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthAccessToken, error) {
	tx := s.db.Begin()

	// 删除过期token
	query := tx.Unscoped().Where("client_id=?", client.ID)
	if user != nil && len([]rune(user.ID)) > 0 {
		query = query.Where("user_id=?", user.ID)
	} else {
		query = query.Where("user_id is NULL")
	}
	if err := query.Where("expires_at<=?", time.Now()).Delete(new(models.OauthAccessToken)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 创建新的token
	accessToken := models.NewOauthAccessToken(client, user, expiresIn, scope)
	if err := tx.Create(accessToken).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	accessToken.Client = client
	accessToken.User = user
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return accessToken, nil
}
