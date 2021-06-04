package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/ariefsn/book-store/api/helper"
	"github.com/ariefsn/book-store/api/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/imroc/req"
)

type AuthController struct {
	BaseController
}

var authUrl = "http://" + os.Getenv("URL_AUTH")

func NewAuthController() *AuthController {
	c := new(AuthController)

	return c
}

func (c *AuthController) BaseUrl() string {
	return authUrl
}

func (c *AuthController) Hi(w http.ResponseWriter, r *http.Request) {
	req := req.New()

	res, err := req.Get(authUrl + "/")

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for login user and get token
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(map[string]interface{}{
			"id":    payload.ID,
			"email": payload.Email,
		}),
	}

	req := req.New()

	res, err := req.Get(authUrl+"/user/me", header)

	if err != nil || res.Response().StatusCode != 200 {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	user := newRes.Data.(map[string]interface{})

	passwordMatch := helper.CheckPasswordHash(payload.Password, user["password"].(string))

	if !passwordMatch {
		render.Render(w, r, helper.ResponseError(http.StatusUnauthorized, errors.New("wrong password")))

		return
	}

	_, token, _ := helper.EncodeJwt(map[string]interface{}{
		"id":    user["id"],
		"email": user["email"],
	})

	render.Render(w, r, helper.ResponseSuccess(token))
}

// Handler for check profile
func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	req := req.New()

	res, err := req.Get(authUrl+"/user/me", header)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for create new user
func (c *AuthController) Create(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	body := req.BodyJSON(&payload)

	req := req.New()

	res, err := req.Post(authUrl+"/user/", header, body)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for register user
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	body := req.BodyJSON(&payload)

	req := req.New()

	res, err := req.Post(authUrl+"/register", body)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for update profile for active user
func (c *AuthController) UpdateMe(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	body := req.BodyJSON(&payload)

	req := req.New()

	res, err := req.Put(authUrl+"/user/me", header, body)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for update user
func (c *AuthController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	body := req.BodyJSON(&payload)

	req := req.New()

	id := chi.URLParam(r, "id")

	res, err := req.Put(authUrl+"/user/"+id, header, body)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for delete user
func (c *AuthController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))
		return
	}

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	body := req.BodyJSON(&payload)

	req := req.New()

	id := chi.URLParam(r, "id")

	res, err := req.Delete(authUrl+"/user/"+id, header, body)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for get all user
func (c *AuthController) All(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	req := req.New()

	res, err := req.Get(authUrl+"/user", header)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for find user
func (c *AuthController) Find(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	req := req.New()

	id := chi.URLParam(r, "id")

	res, err := req.Get(authUrl+"/user/"+id, header)

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		render.Render(w, r, helper.ResponseError(newRes.HTTPStatusCode, errors.New(newRes.Message)))
		return
	}

	render.Render(w, r, helper.Response(&newRes))
}
