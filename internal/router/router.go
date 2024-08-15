package router

import (
	"html/template"
	"net/http"

	"pxgen.io/user/internal/handler"
	"pxgen.io/user/internal/middleware"
	"pxgen.io/user/internal/utils/log"
)

type Router struct {
	userHandler handler.UserHandler
	authHandler handler.AuthHandler
}

func NewRouter(userHandler handler.UserHandler, authHandler handler.AuthHandler) *Router {
	return &Router{userHandler: userHandler, authHandler: authHandler}
}

func (router *Router) SetupRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", middleware.GenerateReqId(middleware.LogCall(router.registerApiVersionOne()))))
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rtpl, err1 := template.ParseFiles("./resources/templates/root.html")
	// etpl, err2 := template.ParseFiles("./templates/error.html")
	if err1 != nil {
		log.Info(err1.Error())
		return
	}
	rtpl.Execute(w, nil)
}

func (router *Router) registerApiVersionOne() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.Handle("/user/", http.StripPrefix("/user", middleware.Auth(router.userRouter())))
	mux.Handle("/auth/", http.StripPrefix("/auth", router.authRouter()))

	return mux

}

func (router *Router) userRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", router.userHandler.CreateUser)
	mux.HandleFunc("PUT /", router.userHandler.UpdateUser)
	mux.HandleFunc("GET /", router.userHandler.GetUser)
	mux.HandleFunc("DELETE /", router.userHandler.DeleteUser)
	return mux
}

func (router *Router) authRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", router.authHandler.Login)
	return mux
}
