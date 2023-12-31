package main

import (
	"cafe/cmd/app/handler"
	boardType3 "cafe/internal/cli/boardType"
	member3 "cafe/internal/cli/member"
	"cafe/internal/controller"
	"cafe/internal/controller/ban"
	"cafe/internal/controller/boardType"
	"cafe/internal/controller/member"
	handler2 "cafe/internal/deco/handler"
	"cafe/internal/repository"
	"cafe/internal/repository/infla"
	"cafe/internal/service"
	ban2 "cafe/internal/service/ban"
	boardType2 "cafe/internal/service/boardType"
	member2 "cafe/internal/service/member"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// 보드 타입
	btH := getBoardTypeHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]}/board-types").Handler(btH)
	r.PathPrefix("/cafes/board-types").Handler(btH)

	// 벤
	bH := getBandHandler()
	r.PathPrefix("/cafes/{cafeId:[0-9]}/ban").Handler(bH)
	r.PathPrefix("/cafes/ban").Handler(bH)

	// 멤버
	mh := getMemberHandler()
	//r.PathPrefix("/cafes/members").Handler(mh)
	//r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(mh)
	r.PathPrefix("/cafes/{cafeId:[0-9]+}/members").Handler(mh)
	r.PathPrefix("/cafes/members").Handler(mh)

	// 카페
	cafeH := getCafeHandler()
	r.PathPrefix("/cafes").Handler(cafeH)

	wrappedRouter := handler2.NewDecoHandler(r, handler2.JwtMiddleware)
	http.ListenAndServe(":8083", wrappedRouter)
}

var boardTypeController = boardType.NewController(boardType2.NewService(boardType3.NewRequester()))
var banController = ban.NewController(ban2.NewService(repository.NewBanRepository(infla.NewDB())))
var cafeController = controller.NewCafeController(service.NewService(repository.NewRepository(infla.NewDB())))
var memberController = member.NewController(member2.NewService(member3.NewRequester()))

func getBoardTypeHandler() http.Handler {
	return handler.NewBoardTypeHandler(boardTypeController, cafeController)
}

func getBandHandler() http.Handler {
	return handler.NewBanHandler(banController, memberController, cafeController)
}
func getMemberHandler() http.Handler {
	return handler.NewMemberHandler(memberController, cafeController)
}

func getCafeHandler() http.Handler {
	return handler.NewCafeHandler(cafeController)
}
