package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ariefsn/book-store/api/controllers"
	"github.com/ariefsn/book-store/api/helper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func main() {
	helper.InitJwt()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	base := controllers.BaseController{}
	auth := controllers.NewAuthController()
	book := controllers.NewBookController()

	r.Get("/", base.Hi)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/", auth.Hi)

		r.Post("/register", auth.Register)
		r.Post("/token", auth.Login)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(helper.TokenAuth()))
			r.Use(helper.Authenticator)

			r.Get("/me", auth.Profile)
			r.Put("/me", auth.UpdateMe)

			r.Get("/user", auth.All)
			r.Get("/user/{id}", auth.Find)
			r.Post("/user", auth.Create)
			r.Put("/user/{id}", auth.UpdateUser)
			r.Delete("/user/{id}", auth.DeleteUser)
		})
	})

	r.Route("/book", func(r chi.Router) {
		r.Get("/hi", book.Hi)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(helper.TokenAuth()))
			r.Use(helper.Authenticator)

			r.Get("/", book.All)
			r.Get("/{id}", book.Find)
			r.Post("/", book.Create)
			r.Put("/{id}", book.UpdateBook)
			r.Delete("/{id}", book.DeleteBook)
		})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, helper.ResponseError(http.StatusMethodNotAllowed, errors.New("method not allowed")))
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, helper.ResponseError(http.StatusNotFound, errors.New("route not found")))
	})

	fmt.Println("\nRegistered Routes")

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("\t%s\t%s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	port := "3001"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("\nServer start on port", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
