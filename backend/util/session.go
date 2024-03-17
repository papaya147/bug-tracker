package util

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func CreateCookieStore(config Config) {
	Store = sessions.NewCookieStore([]byte(config.COOKIE_STORE_SECRET))
	Store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "*",
		MaxAge:   86400 * 30,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		// TODO - uncomment this in production
		// Secure:   true,
	}
}
