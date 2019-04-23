package web

import (
	"fmt"
	"net/http"
	"net/url"
)

func redirectWithQueryString(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s%s", to, getQueryString(query)), http.StatusFound)
}

func redirectWithFragment(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s%s", to, query.Encode()), http.StatusFound)
}
func getQueryString(query url.Values) string {
	encode := query.Encode()
	if len(encode) > 0 {
		encode = fmt.Sprintf("?%s", encode)
	}
	return encode
}

func errorRedirect(w http.ResponseWriter, r *http.Request, redirectURI *url.URL, err, state, responseType string) {
	query := redirectURI.Query()
	query.Set("error", err)
	if state != "" {
		query.Set("state", state)
	}
	if responseType == "code" {
		redirectWithQueryString(redirectURI.String(), query, w, r)
	}
	if responseType == "token" {
		redirectWithFragment(redirectURI.String(), query, w, r)
	}
}
