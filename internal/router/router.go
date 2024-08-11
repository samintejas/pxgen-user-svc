package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"pxgen.io/user/internal/handler"
)

type Router struct {
	userHandler handler.UserHandler
}

func (router *Router) NewRouter(userHandler handler.UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (router *Router) SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/api/v1", router.registerApiVersionOne)

	return r
}

func (router *Router) registerApiVersionOne(r chi.Router) {

	r.Route("/user", func(r chi.Router) {
		r.Post("/", router.userHandler.CreateUser)
		r.Get("/", router.userHandler.ListUsers)
		r.Get("/{id}", router.userHandler.GetUser)
		r.Put("/{id}", router.userHandler.UpdateUser)
		r.Delete("/{id}", router.userHandler.DeleteUser)
	})
}
