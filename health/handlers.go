package health

import (
	"net/http"
	"oauth2-server/util/response"
)

func (s *Service) healthCheck(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Raw("SELECT 1=1").Rows()
	defer rows.Close()
	isHealth := true
	if err != nil {
		isHealth = false
	}
	response.WriteJSON(w, map[string]interface{}{
		"isHealth": isHealth,
	}, 200)
}
