package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ariefsn/book-store/book/helper"
	"github.com/ariefsn/book-store/book/models"
	"github.com/ariefsn/book-store/book/services"
	"github.com/go-chi/render"
)

type BookController struct {
	BaseController
}

func NewBookController() *BookController {
	c := new(BookController)

	return c
}

func (c *BookController) Hi(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, helper.ResponseSuccess("Hi, Welcome to Book Service Version 1"))
}

// Handler for create new book
func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	userId, _ := c.ParseClaims(r)

	payload := models.BookModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	userIdInt, _ := strconv.Atoi(userId)

	user, _, err := services.GetUserByID(r, userIdInt)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	if !user["isAdmin"].(bool) {
		render.Render(w, r, helper.ResponseError(http.StatusUnauthorized, errors.New("role not authorized")))
		return
	}

	id, err := services.CreateBook(&payload)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(id))
}

// Handler for update book
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	payload := models.BookModel{}

	if err := render.Bind(r, &payload); err != nil {
		render.Render(w, r, helper.ResponseError(422, err))

		return
	}

	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	row := services.UpdateBook(id, &payload)

	render.Render(w, r, helper.ResponseSuccess(row))
}

// Handler for delete book
func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	user, err := services.GetBookByID(id)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	row := services.DeleteBook(user)

	render.Render(w, r, helper.ResponseSuccess(row))
}

// Handler for get all books
func (c *BookController) All(w http.ResponseWriter, r *http.Request) {
	userId, _ := c.ParseClaims(r)

	userIdInt, _ := strconv.Atoi(userId)

	admin, _, err := services.GetUserByID(r, userIdInt)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	if !admin["isAdmin"].(bool) {
		render.Render(w, r, helper.ResponseError(http.StatusUnauthorized, errors.New("unauthorized")))
		return
	}

	users, err := services.GetBooks()

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(users))
}

// Handler for find book by id
func (c *BookController) Find(w http.ResponseWriter, r *http.Request) {
	id, code, err := c.ValidateId(r)

	if err != nil {
		render.Render(w, r, helper.ResponseError(code, err))
		return
	}

	user, err := services.GetBookByID(id)

	if err != nil {
		render.Render(w, r, helper.ResponseError(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, helper.ResponseSuccess(user))
}
