package util

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func CreateCookieStore(config Config) {
	Store = sessions.NewCookieStore([]byte(config.COOKIE_STORE_SECRET))
	Store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   86400 * 30,
		HttpOnly: true,
		// TODO - uncomment this in production
		// Secure:   true,
	}
}
