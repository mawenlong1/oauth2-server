package web

import (
	"fmt"
	"net/http"
	"net/url"
)

func redirectWithQueryString(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s%s", to, getQueryString(query)), http.StatusFound)
}
func getQueryString(query url.Values) string {
	encode := query.Encode()
	if len(encode) > 0 {
		encode = fmt.Sprintf("?%s", encode)
	}
	return encode
}
