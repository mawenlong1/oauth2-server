package user

import (
	"net/http"
	"oauth2-server/util/response"
)

var (
	// Superuser ...
	Superuser = "superuser"
	// User ..
	User = "user"
)

func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		response.Error(w, "username 和password不能为空", http.StatusBadRequest)
	}

	if s.oauthService.UserExists(username) {
		response.Error(w, "username已经存在", http.StatusBadRequest)
		return
	}
	_, err := s.oauthService.CreateUser(User, username, password)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.WriteJSON(w, map[string]interface{}{
		"success": true,
	}, 200)
}
