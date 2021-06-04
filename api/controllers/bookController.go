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

type BookController struct {
	BaseController
}

var bookUrl = "http://" + os.Getenv("URL_BOOK")

func NewBookController() *BookController {
	c := new(BookController)

	return c
}

func (c *BookController) BaseUrl() string {
	return bookUrl
}

func (c *BookController) Hi(w http.ResponseWriter, r *http.Request) {
	req := req.New()

	res, err := req.Get(bookUrl + "/")

	if err != nil {
		render.Render(w, r, helper.ResponseError(res.Response().StatusCode, errors.New(res.Response().Status)))
		return
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	render.Render(w, r, helper.Response(&newRes))
}

// Handler for create new book
func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.BookModel{}

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

	res, err := req.Post(bookUrl+"/book/", header, body)

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

// Handler for update book
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.BookModel{}

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

	res, err := req.Put(bookUrl+"/book/"+id, header, body)

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

// Handler for delete book
func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	payload := models.BookModel{}

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

	res, err := req.Delete(bookUrl+"/book/"+id, header, body)

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

// Handler for get all books
func (c *BookController) All(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	req := req.New()

	res, err := req.Get(bookUrl+"/book", header)

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

// Handler for find book
func (c *BookController) Find(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := helper.DecodeJwt(r)

	header := req.Header{
		"Accept": "application/json",
		"Claims": c.BuildClaims(claims),
	}

	req := req.New()

	id := chi.URLParam(r, "id")

	res, err := req.Get(bookUrl+"/book/"+id, header)

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
