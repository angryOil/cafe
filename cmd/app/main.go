package main

import (
	"cafe/cmd/app/handler"
	member3 "cafe/internal/cli/member"
	"cafe/internal/controller"
	"cafe/internal/controller/member"
	handler2 "cafe/internal/deco/handler"
	"cafe/internal/repository"
	"cafe/internal/repository/infla"
	"cafe/internal/service"
	member2 "cafe/internal/service/member"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// 멤버에 관하여
	mh := getMemberHandler()
	//r.PathPrefix("/cafes/members").Handler(mh)
	//r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(mh)
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(mh)
	r.PathPrefix("/cafes/members").Handler(mh)

	cafeH := getCafeHandler()
	r.PathPrefix("/cafes").Handler(cafeH)

	wrappedRouter := handler2.NewDecoHandler(r, handler2.JwtMiddleware)
	http.ListenAndServe(":8083", wrappedRouter)
}

func getMemberHandler() http.Handler {
	return handler.NewMemberHandler(member.NewController(member2.NewService(member3.NewRequester())))
}

func getCafeHandler() http.Handler {
	return handler.NewCafeHandler(controller.NewCafeController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
