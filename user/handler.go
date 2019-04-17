package user

import (
	"net/http"
	"oauth2-server/util/response"
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

	//	TODO校验username是否已经存在并创建（oauth服务实现）
	response.WriteJSON(w, map[string]interface{}{
		"success": true,
	}, 200)
}
