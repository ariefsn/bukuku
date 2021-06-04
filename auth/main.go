package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ariefsn/book-store/auth/controllers"
	"github.com/ariefsn/book-store/auth/helper"
	"github.com/ariefsn/book-store/auth/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	db, err := helper.InitDB()

	if err != nil {
		fmt.Println("[Error]", err.Error())
		return
	}

	err = services.InitService(db)

	if err != nil {
		fmt.Println("[Error]", err.Error())
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	ctr := controllers.NewAuthController()

	r.Get("/", ctr.Hi)
	r.Post("/register", ctr.Register)

	r.Route("/user", func(r chi.Router) {
		r.Get("/", ctr.All)
		r.Get("/{id}", ctr.Find)
		r.Put("/{id}", ctr.UpdateUser)
		r.Delete("/{id}", ctr.DeleteUser)
		r.Get("/me", ctr.Profile)
		r.Put("/me", ctr.UpdateMe)
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

	port := "3002"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("\nServer start on port", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
