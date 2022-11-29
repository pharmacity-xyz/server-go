package controllers

import (
	"fmt"
	"net/http"
)

func AuthorizeAdmin(r *http.Request) error {
	token, err := readCookie(r, COOKIE_TOKEN)
	if err != nil {
		return err
	}

	_, role, err := ParseJWT(token)
	if err != nil {
		return err
	}
	if role != "Admin" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
