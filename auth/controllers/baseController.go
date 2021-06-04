package controllers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/ariefsn/book-store/auth/services"
	"github.com/go-chi/chi/v5"
)

type BaseController struct{}

func (b *BaseController) ParseClaims(r *http.Request) (string, string) {
	claims := r.Header.Get("Claims")

	decoded, _ := base64.StdEncoding.DecodeString(claims)

	split := strings.Split(string(decoded), "*")

	return split[0], split[1]
}

func (b *BaseController) ValidateId(r *http.Request) (int, int, error) {
	_, email := b.ParseClaims(r)

	if chi.URLParam(r, "id") == "" {
		return 0, http.StatusInternalServerError, errors.New("id can't be empty")
	}

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	admin, err := services.GetUserByEmail(email)

	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	if !admin.IsAdmin {
		return 0, http.StatusUnauthorized, errors.New("unauthorized")
	}

	return id, http.StatusOK, nil
}
