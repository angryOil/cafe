package main

import (
	"cafe/cmd/app/handler"
	"cafe/internal/controller"
	handler2 "cafe/internal/deco/handler"
	"cafe/internal/repository"
	"cafe/internal/repository/infla"
	"cafe/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	wrappedMiddleware := handler2.NewDecoHandler(h, handler2.AuthMiddleware)
	r.PathPrefix("/cafes").Handler(wrappedMiddleware)
	http.ListenAndServe(":8083", r)
}

func getHandler() http.Handler {
	return handler.NewCafeHandler(controller.NewCafeController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
