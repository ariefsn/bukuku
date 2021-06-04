package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ariefsn/book-store/auth/helper"
	"github.com/ariefsn/book-store/auth/models"
	"github.com/ariefsn/book-store/auth/services"
	"github.com/go-chi/render"
)

type AuthController struct {
	BaseController
}

func NewAuthController() *AuthController {
	c := new(AuthController)

	return c
}

func (c *AuthController) Hi(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, helper.ResponseSuccess("Hi, Welcome to Auth Service Version 1"))
}

// Handler for register user
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	// check email
	checkUser, _ := services.GetUserByEmail(payload.Email)

	if payload.Email == checkUser.Email {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, errors.New("email registered")))
		return
	}

	id, err := services.CreateUser(&payload)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(id))
}

// Handler for create new user
func (c *AuthController) Create(w http.ResponseWriter, r *http.Request) {
	_, email := c.ParseClaims(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	user, err := services.GetUserByEmail(email)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	if !user.IsAdmin {
		render.Render(w, r, helper.ResponseError(http.StatusUnauthorized, errors.New("role not authorized")))
		return
	}

	// check email
	checkUser, _ := services.GetUserByEmail(payload.Email)

	if payload.Email == checkUser.Email {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, errors.New("email registered")))
		return
	}

	id, err := services.CreateUser(&payload)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(id))
}

// Handler for check profile
func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	_, email := c.ParseClaims(r)

	user, err := services.GetUserByEmail(email)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))

		return
	}

	render.Render(w, r, helper.ResponseSuccess(user))
}

// Handler for update active user
func (c *AuthController) UpdateMe(w http.ResponseWriter, r *http.Request) {
	sId, _ := c.ParseClaims(r)

	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	id, _ := strconv.Atoi(sId)

	row := services.UpdateUser(id, &payload)

	render.Render(w, r, helper.ResponseSuccess(row))
}

// Handler for update user
func (c *AuthController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	payload := models.UserModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	row := services.UpdateUser(id, &payload)

	render.Render(w, r, helper.ResponseSuccess(row))
}

// Handler for delete user
func (c *AuthController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	user, err := services.GetUserByID(id)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	row := services.DeleteUser(user)

	render.Render(w, r, helper.ResponseSuccess(row))
}

// Handler for get all users
func (c *AuthController) All(w http.ResponseWriter, r *http.Request) {
	_, email := c.ParseClaims(r)

	admin, err := services.GetUserByEmail(email)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	if !admin.IsAdmin {
		render.Render(w, r, helper.ResponseError(http.StatusUnauthorized, errors.New("unauthorized")))
		return
	}

	users, err := services.GetUsers()

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(users))
}

// Handler for find user
func (c *AuthController) Find(w http.ResponseWriter, r *http.Request) {
	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	user, err := services.GetUserByID(id)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(user))
}
