package router

import (
	"fmt"
	"net/http"

	"pxgen.io/user/internal/handler"
	"pxgen.io/user/internal/middleware"
	"pxgen.io/user/internal/utils"
)

type Router struct {
	userHandler handler.UserHandler
}

func NewRouter(userHandler handler.UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (router *Router) SetupRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", middleware.GenerateReqId(middleware.LogCall(router.registerApiVersionOne()))))
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid endpoint"))
}

func (router *Router) registerApiVersionOne() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.Handle("/user/", http.StripPrefix("/user", router.userRouter()))
	mux.Handle("/auth/", http.StripPrefix("/auth", authRouter()))

	return mux

}

func (router *Router) userRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", router.userHandler.CreateUser)
	mux.HandleFunc("PUT /{id}", router.userHandler.UpdateUser)
	mux.HandleFunc("GET /{id}", router.userHandler.GetUser)
	mux.HandleFunc("DELETE /{id}", router.userHandler.DeleteUser)
	return mux
}

func authRouter() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
