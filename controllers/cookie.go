package controllers

import (
	"fmt"
	"net/http"
	"time"
)

const (
	COOKIE_TOKEN = "token"
)

func newCookie(name, value string, expirationTime time.Time) *http.Cookie {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expirationTime,
	}
	return &cookie
}

func setCookie(w http.ResponseWriter, name, value string, expirationTime time.Time) {
	cookie := newCookie(name, value, expirationTime)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return c.Value, nil
}

func deleteCookie(w http.ResponseWriter, name string) {
	cookie := newCookie(name, "", time.Now())
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
