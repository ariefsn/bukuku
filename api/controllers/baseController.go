package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/ariefsn/book-store/api/helper"
	"github.com/go-chi/render"
)

type BaseController struct{}

func (c *BaseController) Hi(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, helper.ResponseSuccess("Hi, Welcome to API Gateway Version 1"))
}

func (c *BaseController) BuildClaims(claims map[string]interface{}) string {
	id := claims["id"]
	email := claims["email"]

	claimsString := fmt.Sprintf("%v*%s", id, email)
	data := []byte(claimsString)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded
}
